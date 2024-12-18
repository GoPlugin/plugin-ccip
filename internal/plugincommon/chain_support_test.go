package plugincommon

import (
	"fmt"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/require"

	"github.com/goplugin/plugin-libocr/commontypes"
	"github.com/goplugin/plugin-libocr/ragep2p/types"

	"github.com/goplugin/plugin-common/pkg/logger"

	reader2 "github.com/goplugin/plugin-ccip/internal/reader"
	"github.com/goplugin/plugin-ccip/mocks/internal_/reader"
	cciptypes "github.com/goplugin/plugin-ccip/pkg/types/ccipocr3"
)

const (
	srcChainA = cciptypes.ChainSelector(0xa)
	srcChainB = cciptypes.ChainSelector(0xb)
	srcChainC = cciptypes.ChainSelector(0xc)
	dstChain  = cciptypes.ChainSelector(0xde)
)

func TestCCIPChainSupport_KnownSourceChainsSlice(t *testing.T) {
	lggr := logger.Test(t)
	homeChainReader := reader.NewMockHomeChain(t)
	cs := &CCIPChainSupport{
		lggr:      lggr,
		homeChain: homeChainReader,
		destChain: dstChain,
	}

	t.Run("happy path", func(t *testing.T) {
		homeChainReader.EXPECT().GetKnownCCIPChains().
			Return(mapset.NewSet(srcChainA, srcChainB, srcChainC), nil).Once()
		knownSourceChains, err := cs.KnownSourceChainsSlice()
		require.NoError(t, err)
		require.Equal(t, []cciptypes.ChainSelector{srcChainA, srcChainB, srcChainC}, knownSourceChains)
	})

	t.Run("error path", func(t *testing.T) {
		homeChainReader.EXPECT().GetKnownCCIPChains().Return(nil, fmt.Errorf("some err")).Once()
		_, err := cs.KnownSourceChainsSlice()
		require.Error(t, err)
	})
}

func TestCCIPChainSupport_SupportedChains(t *testing.T) {
	lggr := logger.Test(t)
	homeChainReader := reader.NewMockHomeChain(t)
	cs := &CCIPChainSupport{
		lggr:            lggr,
		homeChain:       homeChainReader,
		oracleIDToP2PID: map[commontypes.OracleID]types.PeerID{1: [32]byte{1}},
	}

	t.Run("happy path", func(t *testing.T) {
		exp := mapset.NewSet(srcChainA, srcChainB, srcChainC)
		homeChainReader.EXPECT().GetSupportedChainsForPeer(types.PeerID{1}).Return(exp, nil).Once()
		supportedChains, err := cs.SupportedChains(1)
		require.NoError(t, err)
		require.True(t, exp.Equal(supportedChains))
	})

	t.Run("oracle not found", func(t *testing.T) {
		_, err := cs.SupportedChains(2)
		require.Error(t, err)
	})

	t.Run("home chain reader error", func(t *testing.T) {
		homeChainReader.EXPECT().GetSupportedChainsForPeer(types.PeerID{1}).
			Return(nil, fmt.Errorf("some err")).Once()
		_, err := cs.SupportedChains(1)
		require.Error(t, err)
	})
}

func TestCCIPChainSupport_SupportsDestChain(t *testing.T) {
	lggr := logger.Test(t)
	homeChainReader := reader.NewMockHomeChain(t)
	cs := &CCIPChainSupport{
		lggr:            lggr,
		homeChain:       homeChainReader,
		destChain:       dstChain,
		oracleIDToP2PID: map[commontypes.OracleID]types.PeerID{1: [32]byte{1}},
	}

	t.Run("happy path", func(t *testing.T) {
		supportedNodes := mapset.NewSet(types.PeerID{1})
		homeChainReader.EXPECT().GetChainConfig(dstChain).
			Return(reader2.ChainConfig{SupportedNodes: supportedNodes}, nil).Once()
		supports, err := cs.SupportsDestChain(1)
		require.NoError(t, err)
		require.True(t, supports)
	})

	t.Run("oracle not found error", func(t *testing.T) {
		supportedNodes := mapset.NewSet(types.PeerID{1})
		homeChainReader.EXPECT().GetChainConfig(dstChain).
			Return(reader2.ChainConfig{SupportedNodes: supportedNodes}, nil).Once()
		_, err := cs.SupportsDestChain(2)
		require.Error(t, err)
	})

	t.Run("not supported", func(t *testing.T) {
		supportedNodes := mapset.NewSet(types.PeerID{2})
		homeChainReader.EXPECT().GetChainConfig(dstChain).
			Return(reader2.ChainConfig{SupportedNodes: supportedNodes}, nil).Once()
		supports, err := cs.SupportsDestChain(1)
		require.NoError(t, err)
		require.False(t, supports)
	})
}
