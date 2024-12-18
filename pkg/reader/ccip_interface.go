package reader

import (
	"context"
	"errors"
	"time"

	"github.com/goplugin/plugin-common/pkg/logger"
	"github.com/goplugin/plugin-common/pkg/types"

	rmntypes "github.com/goplugin/plugin-ccip/commit/merkleroot/rmn/types"
	"github.com/goplugin/plugin-ccip/internal/plugintypes"

	"github.com/goplugin/plugin-ccip/pkg/contractreader"
	cciptypes "github.com/goplugin/plugin-ccip/pkg/types/ccipocr3"
	plugintypes2 "github.com/goplugin/plugin-ccip/plugintypes"
)

var (
	ErrContractReaderNotFound = errors.New("contract reader not found")
	ErrContractWriterNotFound = errors.New("contract writer not found")
)

// ContractAddresses is a map of contract names across all chain selectors and their address.
// Currently only one contract per chain per name is supported.
type ContractAddresses map[string]map[cciptypes.ChainSelector][]byte

func (ca ContractAddresses) Append(contract string, chain cciptypes.ChainSelector, address []byte) ContractAddresses {
	resp := ca
	if resp == nil {
		resp = make(ContractAddresses)
	}
	if resp[contract] == nil {
		resp[contract] = make(map[cciptypes.ChainSelector][]byte)
	}
	resp[contract][chain] = address
	return resp
}

func NewCCIPChainReader(
	ctx context.Context,
	lggr logger.Logger,
	contractReaders map[cciptypes.ChainSelector]contractreader.ContractReaderFacade,
	contractWriters map[cciptypes.ChainSelector]types.ChainWriter,
	destChain cciptypes.ChainSelector,
	offrampAddress []byte,
) CCIPReader {
	return newCCIPChainReaderInternal(
		ctx,
		lggr,
		contractReaders,
		contractWriters,
		destChain,
		offrampAddress,
	)
}

// NewCCIPReaderWithExtendedContractReaders can be used when you want to directly provide contractreader.Extended
func NewCCIPReaderWithExtendedContractReaders(
	ctx context.Context,
	lggr logger.Logger,
	contractReaders map[cciptypes.ChainSelector]contractreader.Extended,
	contractWriters map[cciptypes.ChainSelector]types.ChainWriter,
	destChain cciptypes.ChainSelector,
	offrampAddress []byte,
) CCIPReader {
	cr := newCCIPChainReaderInternal(ctx, lggr, nil, contractWriters, destChain, offrampAddress)
	for ch, extendedCr := range contractReaders {
		cr.WithExtendedContractReader(ch, extendedCr)
	}
	return cr
}

type CCIPReader interface {
	// CommitReportsGTETimestamp reads the requested chain starting at a given timestamp
	// and finds all ReportAccepted up to the provided limit.
	CommitReportsGTETimestamp(
		ctx context.Context,
		dest cciptypes.ChainSelector,
		ts time.Time,
		limit int,
	) ([]plugintypes2.CommitPluginReportWithMeta, error)

	// ExecutedMessageRanges reads the destination chain and finds which messages are executed.
	// A slice of sequence number ranges is returned to express which messages are executed.
	ExecutedMessageRanges(
		ctx context.Context,
		source, dest cciptypes.ChainSelector,
		seqNumRange cciptypes.SeqNumRange,
	) ([]cciptypes.SeqNumRange, error)

	// MsgsBetweenSeqNums reads the provided chains.
	// Finds and returns ccip messages submitted between the provided sequence numbers.
	// Messages are sorted ascending based on their timestamp and limited up to the provided limit.
	MsgsBetweenSeqNums(
		ctx context.Context,
		chain cciptypes.ChainSelector,
		seqNumRange cciptypes.SeqNumRange,
	) ([]cciptypes.Message, error)

	// GetExpectedNextSequenceNumber returns the next sequence number to be used
	// in the onramp.
	GetExpectedNextSequenceNumber(
		ctx context.Context,
		sourceChainSelector, destChainSelector cciptypes.ChainSelector,
	) (cciptypes.SeqNum, error)

	// NextSeqNum reads the destination chain.
	// Returns the next expected sequence number for each one of the provided chains.
	// TODO: if destination was a parameter, this could be a capability reused across plugin instances.
	NextSeqNum(ctx context.Context, chains []cciptypes.ChainSelector) (seqNum []cciptypes.SeqNum, err error)

	// GetContractAddress returns the contract address that is registered for the provided contract name and chain.
	GetContractAddress(contractName string, chain cciptypes.ChainSelector) ([]byte, error)

	// Nonces fetches all nonces for the provided selector/address pairs. Addresses are a string encoded raw address,
	// it must be encoding according to the destination chain requirements with typeconv.AddressBytesToString.
	Nonces(
		ctx context.Context,
		source, dest cciptypes.ChainSelector,
		addresses []string,
	) (map[string]uint64, error)

	// GetAvailableChainsFeeComponents Reads all fee components for known chains (chains that have chain writer defined)
	GetAvailableChainsFeeComponents(ctx context.Context) map[cciptypes.ChainSelector]types.ChainFeeComponents

	// GetWrappedNativeTokenPriceUSD Gets the wrapped native token price in USD for the provided chains.
	GetWrappedNativeTokenPriceUSD(
		ctx context.Context,
		selectors []cciptypes.ChainSelector,
	) map[cciptypes.ChainSelector]cciptypes.BigInt

	// GetChainFeePriceUpdate Gets latest chain fee price update for the provided chains.
	GetChainFeePriceUpdate(
		ctx context.Context,
		selectors []cciptypes.ChainSelector,
	) map[cciptypes.ChainSelector]plugintypes.TimestampedBig

	GetRMNRemoteConfig(
		ctx context.Context,
		destChainSelector cciptypes.ChainSelector,
	) (rmntypes.RemoteConfig, error)

	// DiscoverContracts reads from all available contract readers to discover contract addresses.
	DiscoverContracts(ctx context.Context) (ContractAddresses, error)

	// LinkPriceUSD gets the PLI price in 1e-18 USDs from the FeeQuoter contract on the destination chain.
	// For example, if the price is 1 PLI = 10 USD, this function will return 10e18 (10 * 1e18). You can think of this
	// function returning the price of PLI not in USD, but in a small denomination of USD, similar to returning
	// the price of ETH not in ETH but in wei (1e-18 ETH).
	LinkPriceUSD(ctx context.Context) (cciptypes.BigInt, error)

	// Sync can be used to perform frequent syncing operations inside the reader implementation.
	// Returns a bool indicating whether something was updated.
	Sync(ctx context.Context, contracts ContractAddresses) error
}
