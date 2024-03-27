package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/Stride-Labs/stride/v21/x/stakeibc/keeper"
	"github.com/Stride-Labs/stride/v21/x/stakeibc/types"
)

func SimulateMsgRestoreInterchainAccount(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRestoreInterchainAccount{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RestoreInterchainAccount simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RestoreInterchainAccount simulation not implemented"), nil, nil
	}
}
