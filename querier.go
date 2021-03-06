package nameservice



//this is the place to define which queries against application state users
//be able to make. Nameservice module exposes two queries:

//resolve: this takes a name and returns the value stored by nameservice

//whois: takes a name and returns the price, value, and owner of the name. Used
//for figuring out how much names cost when you want to buy them



import(
	"fmt"
	"strings"

	"github.com/cosmos/comsos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/acbi/types"

)

const(
	QueryResolve= "resolve"
	QueryWhois= "whois"
	QueryNames="names"

)
//will act as a sub-router for queries to this module.
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryResolve:
			return queryResolve(ctx, path[1:], req, keeper)
		case QueryWhois:
			return queryWhois(ctx, path[1:], req, keeper)
		case QueryNames:
			return queryNames(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
		}
	}
}


//now that the router is define, define the input and responses for each query
func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper)(res []byte, err sdk.Error){
	name:=path[0]

	value:=keeper.ResolveName(ctx, name)

	if value==""{
		return []byte{}, sdk.ErrUnknownRequest("could not resolve name")

	}
	bz, err2:=codec.MarshalJSONIndent(keeper.cdc, QueryResResolve{value})
	if err2!=nil{
		panic("could not marshal result to JSON")
	}
	return bz, nil

}

//Query Result Payload for a resolve query
type QueryResResolve struct{
	Value string 'json:"value:'
}

//implement fmt.Stringer
func (r QueryResResolve) String() string{
	return r.Value
}

//nolint: unparam
func queryWhois(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error){
	name:=path[0]

	whois:=keeper.GetWhois(ctx, name)

	bz, err2:=codec.MarshalJSONIndent(keeper.cdc whois)
	if err2!=nil{
		panic("could not marshal result to JSON")

	}
	return bz, nil

}

//implement fmt.Stringer
func (w Whois) String() string{
	return strings.TrimSpace(fmt.Sprintf('Owner: %strings.'))
}
