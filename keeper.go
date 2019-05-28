package nameservice

//The Main core of the Cosmos SDK module. It is what handles interactions with the store,
//has references to other keepers for cross-module interactions,
//and contains most of the core functionality of a module

//Three cosmos-sdk packages are imported: codec (for Amino encoding format), bank, and types

//

import (
	"github.com/cosmos.cosmos-sdk/codec"  //provides tools to work with the Cosmos encoding format
	"github.com/cosmos/cosmos-sdk/x/bank" //controls accounts and coin transfers

	sdk "github.com/cosmos/cosmos-sdk/types" //contains commonly used types throughout SDK
)

//Keeper maintains the link to data storage and exposes getter/setter
//methods for various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper //reference to the Keeper from the "bank" module.
	//including it allows code in this module to call functions
	//from the "bank" module.

	storeKey sdk.Storekey //unexposed key to access store from sdk.Context

	cdc *codec.Codec //The wire codec for binary encoding/decoding. It's a pointer
	//to the codec that is used by Amino t encode and decode
	//binary structs
}

//sets the entire Whois metadata struct for a name
func (k Keeper) SetWhois(ctx sdk.Context, name string, whois Whois) {
	if whois.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storekey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(whois))

}

//***************time to add methods to interact with the stores through the Keeper ************************

//Gets the whole Whois metadata struct for a name
func (k Keeper) GetWhois(ctx sdk.Context, name string) Whois {

	store := ctx.KVStore(k.storekey) //access store using this storekey
	if !store.Has([]byte(name)) {    //if a name doesn't exist in store, make new one
		return NewWhois() //which has minimumPrice initialized in it
	}
	bz := store.Get([]byte(name))
	var whois Whois
	k.cdc.MustUnmarshalBinaryBare(bz, &whois) //Unmarshals the byteslice back into
	// a Whois struct which we return
	return whois
}

//List of functions for getting specific parameters from store based on the name

//ResolveName- returns the string that the name resolves to
func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
	return k.GetWhois(ctx, name).Value
}

//SetName- sets the value string that a name resolves to
func (k Keeper) SetName(ctx sdk.Context, name string, value string) {

	whois := k.GetWhois(ctx, name)
	whois.Value = value
	k.SetWhois(ctx, name, whois)
}

//gets the current price of a name. If price doesn't exist yet, set to 1nametoken
func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	return k.GetWhois(ctx, name).Price
}

//SetPrice- sets the current price of a name
func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	whois := k.GetWhois(ctx, name)
	whois.Price = price
	k.SetWhois(ctx, name, whois)
}

//get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

// constructor for the nameservice Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}
