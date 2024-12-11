package execute

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/goplugin/plugin-common/pkg/utils/tests"

	"github.com/goplugin/plugin-ccip/execute/exectypes"
	"github.com/goplugin/plugin-ccip/internal/mocks/inmem"
	"github.com/goplugin/plugin-ccip/internal/plugintypes"
	cciptypes "github.com/goplugin/plugin-ccip/pkg/types/ccipocr3"
)

func TestPlugin(t *testing.T) {
	ctx := tests.Context(t)

	srcSelector := cciptypes.ChainSelector(1)
	dstSelector := cciptypes.ChainSelector(2)

	messages := []inmem.MessagesWithMetadata{
		makeMsg(100, srcSelector, dstSelector, true),
		makeMsg(101, srcSelector, dstSelector, true),
		makeMsg(102, srcSelector, dstSelector, false),
		makeMsg(103, srcSelector, dstSelector, false),
		makeMsg(104, srcSelector, dstSelector, false),
		makeMsg(105, srcSelector, dstSelector, false),
	}

	intTest := SetupSimpleTest(t, srcSelector, dstSelector)
	intTest.WithMessages(messages, 1000, time.Now().Add(-4*time.Hour))
	runner := intTest.Start()
	defer intTest.Close()

	// Contract Discovery round.
	outcome := runner.MustRunRound(ctx, t)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// Round 1 - Get Commit Reports
	// One pending commit report only.
	// Two of the messages are executed which should be indicated in the Outcome.
	outcome = runner.MustRunRound(ctx, t)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.PendingCommitReports, 1)
	require.ElementsMatch(t, outcome.PendingCommitReports[0].ExecutedMessages, []cciptypes.SeqNum{100, 101})

	// Round 2 - Get Messages
	// Messages now attached to the pending commit.
	outcome = runner.MustRunRound(ctx, t)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.PendingCommitReports, 1)

	// Round 3 - Filter
	// An execute report with the following messages executed: 102, 103, 104, 105.
	outcome = runner.MustRunRound(ctx, t)
	require.Len(t, outcome.Report.ChainReports, 1)
	sequenceNumbers := extractSequenceNumbers(outcome.Report.ChainReports[0].Messages)
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{102, 103, 104, 105})
}

func Test_ExcludingCostlyMessages(t *testing.T) {
	ctx := tests.Context(t)

	srcSelector := cciptypes.ChainSelector(1)
	dstSelector := cciptypes.ChainSelector(2)

	messages := []inmem.MessagesWithMetadata{
		makeMsg(100, srcSelector, dstSelector, false, withFeeValueJuels(100)),
		makeMsg(101, srcSelector, dstSelector, false, withFeeValueJuels(200)),
		makeMsg(102, srcSelector, dstSelector, false, withFeeValueJuels(300)),
	}

	messageTimestamp := time.Now().Add(-4 * time.Hour)
	tm := timeMachine{now: messageTimestamp}

	intTest := SetupSimpleTest(t, srcSelector, dstSelector)
	intTest.WithMessages(messages, 1000, messageTimestamp)
	intTest.WithCustomFeeBoosting(1.0, tm.Now, map[cciptypes.Bytes32]plugintypes.USD18{
		messages[0].Header.MessageID: plugintypes.NewUSD18(40000),
		messages[1].Header.MessageID: plugintypes.NewUSD18(200000),
		messages[2].Header.MessageID: plugintypes.NewUSD18(200000),
	})

	runner := intTest.Start()
	defer intTest.Close()

	outcome := runner.MustRunRound(ctx, t)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// First outcome is empty - all messages are too expensive to be executed
	// Message1 cost=40000,  fee=10000
	// Message2 cost=200000, fee=20000
	// Message3 cost=200000, fee=30000
	for i := 0; i < 3; i++ {
		outcome = runner.MustRunRound(ctx, t)
	}
	require.Len(t, outcome.Report.ChainReports, 0)

	// 4 hours later, we agree to pay higher fee, but only for the first message
	// Message1 cost=40000,  fee=50000 boosted original_fee * (1 + 4*1.0),
	// Message2 cost=200000, fee=20000
	// Message3 cost=200000, fee=30000
	tm.SetNow(time.Now())
	for i := 0; i < 3; i++ {
		outcome = runner.MustRunRound(ctx, t)
	}
	sequenceNumbers := extractSequenceNumbers(outcome.Report.ChainReports[0].Messages)
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{100})

	// Second message execution cost drops, it should be included in the outcome
	// Message1 cost=40000,  fee=50000   boosted original_fee * (1 + 4*1.0),
	// Message2 cost=40000,  fee=100000
	// Message3 cost=200000, fee=150000
	intTest.UpdateExecutionCost(messages[1].Header.MessageID, 40000)
	for i := 0; i < 3; i++ {
		outcome = runner.MustRunRound(ctx, t)
	}
	sequenceNumbers = extractSequenceNumbers(outcome.Report.ChainReports[0].Messages)
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{100, 101})

	// 3 hours in the future, we agree to pay higher fee for the third message (7 hours since the message was sent)
	// Message1 cost=40000,  fee=80000  boosted 10000 * (1 + 7*1.0),
	// Message2 cost=40000,  fee=160000
	// Message3 cost=200000, fee=240000
	tm.SetNow(time.Now().Add(3 * time.Hour))
	for i := 0; i < 3; i++ {
		outcome = runner.MustRunRound(ctx, t)
	}
	sequenceNumbers = extractSequenceNumbers(outcome.Report.ChainReports[0].Messages)
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{100, 101, 102})
}