package nameservice

//This is a struct that is going to hold all the metadata of a name
//this date includes:

//Value: the value that a name resolves to. This is just an arbitrary string, but
//in the future it can be modified to a DNS Zone File, Ip address, etc.

//Owner: the owner of the name

//Price: the price you will need to pay in order to buy the name

import(
	"sdk github.com/cosmos/cosmos-sdk/types"
)

//Whois contains all the metadata of a name
type Whois struct{
	Value string			'json:"value:"'
	Owner sdk.AccAdress		'json:"owner:"'
	Price sdk.Coins			'json:"price:"'
}

//Initial Starting Price for a name that was never previously owned
var minNamePrice=sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

//return a new Whois with the minprice as the price
func NewWhois() Whois{
	return Whois{
		Price: MinNamePrice,
	}
}



