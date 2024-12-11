package gas

import "github.com/goplugin/plugin-ccip/pkg/types/ccipocr3"

// EstimateProvider is used to estimate the gas cost of a message or a merkle tree.
// TODO: Move to pkg/types/ccipocr3 or remove.
type EstimateProvider interface {
	CalculateMerkleTreeGas(numRequests int) uint64
	CalculateMessageMaxGas(msg ccipocr3.Message) uint64
}
