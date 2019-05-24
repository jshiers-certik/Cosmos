package nameservice

//The Main core of the Cosmos SDK module. It is what handles interactions with the store,
//has references to other keepers for cross-module interactions,
//and contains most of the core functionality of a module

import(
	"github.com/cosmos.cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"

)

//Keeper maintains the link to data storage and exposes getter/setter
//methods for various parts of the state machine
type Keeper struct{
	coinKeeper bank.Keeper
	storeKey sdk.Storekey //unexposed key to access store from sdk.Context
	cdc *codec.Codec //The wire codec for binary encoding/encoding
}
