package reader

import (
	"context"
	"fmt"

	"github.com/goplugin/plugin-common/pkg/logger"
	"github.com/goplugin/plugin-common/pkg/types"

	typeconv "github.com/goplugin/plugin-ccip/internal/libs/typeconv"
	"github.com/goplugin/plugin-ccip/pkg/contractreader"
	cciptypes "github.com/goplugin/plugin-ccip/pkg/types/ccipocr3"
)

// bindable is a helper interface to represent all different types of contract readers.
// We expose per-reader type functions, but then cast to this common interface to avoid code duplication.
type bindable interface {
	Bind(ctx context.Context, bindings []types.BoundContract) error
}

func bindExtendedReaderContract(
	ctx context.Context,
	lggr logger.Logger,
	readers map[cciptypes.ChainSelector]contractreader.Extended,
	chainSel cciptypes.ChainSelector,
	contractName string,
	address []byte,
) (types.BoundContract, error) {
	casted := make(map[cciptypes.ChainSelector]bindable, len(readers))
	for k, v := range readers {
		casted[k] = v
	}
	return bindReaderContract(ctx, lggr, casted, chainSel, contractName, address)
}

func bindFacadeReaderContract(
	ctx context.Context,
	lggr logger.Logger,
	readers map[cciptypes.ChainSelector]contractreader.ContractReaderFacade,
	chainSel cciptypes.ChainSelector,
	contractName string,
	address []byte,
) (types.BoundContract, error) {
	casted := make(map[cciptypes.ChainSelector]bindable, len(readers))
	for k, v := range readers {
		casted[k] = v
	}
	return bindReaderContract(ctx, lggr, casted, chainSel, contractName, address)
}

// bindReaderContract is a generic helper for binding contracts to readers, the addresses input is the same object
// returned by DiscoverContracts.
//
// No error is returned if contractName is not found in the contracts. This allows calling the function before all
// contracts are discovered.
func bindReaderContract(
	ctx context.Context,
	lggr logger.Logger,
	readers map[cciptypes.ChainSelector]bindable,
	chainSel cciptypes.ChainSelector,
	contractName string,
	address []byte,
) (types.BoundContract, error) {
	var empty types.BoundContract

	if err := validateReaderExistence(readers, chainSel); err != nil {
		return empty, fmt.Errorf("validate reader existence: %w", err)
	}

	encAddress := typeconv.AddressBytesToString(address, uint64(chainSel))
	contract := types.BoundContract{
		Address: encAddress,
		Name:    contractName,
	}

	lggr.Debugw("Binding contract",
		"chainSel", chainSel,
		"contractName", contractName,
		"address", encAddress,
	)
	// Bind the contract address to the reader.
	// If the same address exists -> no-op
	// If the address is changed -> updates the address, overwrites the existing one
	// If the contract not bound -> binds to the new address
	if err := readers[chainSel].Bind(ctx, []types.BoundContract{contract}); err != nil {
		return empty, fmt.Errorf("unable to bind %s %s for chain %d: %w", contractName, encAddress, chainSel, err)
	}

	return contract, nil
}

func validateExtendedReaderExistence(
	readers map[cciptypes.ChainSelector]contractreader.Extended,
	chains ...cciptypes.ChainSelector,
) error {
	casted := make(map[cciptypes.ChainSelector]bindable, len(readers))
	for k, v := range readers {
		casted[k] = v
	}
	return validateReaderExistence(casted, chains...)
}

func validateReaderExistence(
	readers map[cciptypes.ChainSelector]bindable,
	chains ...cciptypes.ChainSelector,
) error {
	for _, ch := range chains {
		_, exists := readers[ch]
		if !exists {
			return fmt.Errorf("chain %d: %w", ch, ErrContractReaderNotFound)
		}
	}
	return nil
}