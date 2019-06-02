package leverage

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cicizeo/hilo/x/leverage/keeper"
	"github.com/cicizeo/hilo/x/leverage/types"
)

// InitGenesis initializes the x/leverage module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	for _, asset := range genState.Assets {
		k.SetAsset(ctx, asset)
	}
}

// ExportGenesis returns the x/leverage module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	return genesis
}
