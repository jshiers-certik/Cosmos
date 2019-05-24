package nameservice

import(
	"sdk github.com/cosmos/cosmos-sdk/types"

)

//Whois contains all the metadata of a name
type Whois struct{
	Value string			'json:"value"'
	Owner sdk.AccAdress		'json:"owner:"'
	Price sdk.Coins			'json:"price:"'
}

//Initial Starting Price for a name that was never previously owned
var minNamePrice=sdk.Coins{
	sdk.NewInt64Coin("nametoken", 1)
}

//return a new Whois with the minprice as the price
func NewWhois() Whois{
	return Whois{
		Price: MinNamePrice,
	}
}

