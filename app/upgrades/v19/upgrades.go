package v19

import (
	errorsmod "cosmossdk.io/errors"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

const (
	UpgradeName = "v19"

	WasmAdmin = "stride159smvptpq6evq0x6jmca6t8y7j8xmwj6kxapyh"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v19
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	cdc codec.Codec,
	wasmKeeper wasmkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("Starting upgrade v19...")

		// Run module migrations first to add wasm to the store
		ctx.Logger().Info("Running module migrations...")
		newVm, err := mm.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return newVm, err
		}

		if err := SetWasmPermissions(ctx, wasmKeeper); err != nil {
			return newVm, errorsmod.Wrapf(err, "unable to set wasm permissions")
		}

		return newVm, nil
	}
}

// Update wasm params so that contracts can only be uploaded through governance
func SetWasmPermissions(ctx sdk.Context, wk wasmkeeper.Keeper) error {
	wasmParams := wk.GetParams(ctx)
	wasmParams.CodeUploadAccess = wasmtypes.AccessConfig{
		Permission: wasmtypes.AccessTypeAnyOfAddresses,
		Addresses:  []string{WasmAdmin},
	}
	wasmParams.InstantiateDefaultPermission = wasmtypes.AccessTypeNobody
	if err := wk.SetParams(ctx, wasmParams); err != nil {
		return err
	}
	return nil
}
