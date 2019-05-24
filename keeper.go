package nameservice

//The Main core of the Cosmos SDK module. It is what handles interactions with the store,
//has references to other keepers for cross-module interactions,
//and contains most of the core functionality of a module

import(
	"github.com/cosmos.cosmos-sdk/codec"  //provides tools to work with the Cosmos encoding format
	"github.com/cosmos/cosmos-sdk/x/bank"  //controls accounts and coin transfers

	sdk "github.com/cosmos/cosmos-sdk/types" //contains commonly used types throughout SDK


)

//Keeper maintains the link to data storage and exposes getter/setter
//methods for various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper //reference to the Keeper from the "bank" module.
	//including it allows code in this module to call functions
	//from the "bank" module.
	storeKey sdk.Storekey //unexposed key to access store from sdk.Context
	cdc      *codec.Codec //The wire codec for binary encoding/decoding. It's a pointer
	 						//to the codec that is used by Amino t encode and decode
	 						//binary structs
}

//sets the entire Whois metadata struct for a name
func (k Keeper) SetWhois(ctx sdk.Context, name string, whois Whois){
	if whois.Owner.Empty(){
		return
	}
	store:=ctx.KVStore(k.storekey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(whois))

}

func (k Keeper) GetWhois(ctx sdk.Context, name string) Whois {

	store:=ctx.KVStore(k.storekey)
	if !store.Has([]byte(name)) {
		return NewWhois()
	}
	bz:=store.Get([]byte(name))
	var whois Whois
	k.cdc.MustUnmarshalBinaryBare(bz, &whois)
	return whois
}










}








