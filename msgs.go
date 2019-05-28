package nameservice

import(
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"

)

//MsgSetName defines a SetName message
type MsgSetName struct{
	Name string
	Value string
	Owner sdk.AccAddress
}

//is a constructor function for MsgSetName
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName{
	return MsgSetName{
		Name: name,
		Value: value,
		Owner: owner,

	}

}


//MsgSetName has 3 attributes needed to set the value for a name:
//name- the name trying to be set
//value- what the name resolves to
//owner- the owner of that name

//Route should return the name of the module
func (msg MsgSetName) Route() string {return "nameservice"}

//type should return the action
func (msg MsgSetName) Type() string{return "set_name"}

//the above functions are used by the SDK to route Msgs to the proper module for handling
//they also add human readable names to the database tags used for indexing


//ValidateBasic runs stateless checks on the message.
func(msg MsgSetName) ValidateBasic() sdk.Error{
	if msg.Owner.Empty(){
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name)==0 || len(msg.Value)==0 {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")

	}
	return nil
}

//encodes the message for signing. Defines how the Msg gets encoded, im most cases
//this means marshal to sorted JSON. The output should not be modified
func (msg MsgSetName) GetSignBytes() []byte {
	b, err:=json.Marshal(msg)
	if err!=nil{
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

//defines whose signature is required on a Tx in order for it to be valid. In this case,
//requires that the Owner signs the transaction when trying to reset what the name
//points to
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}



type  MsgBuyName struct{
	Name string
	Bid sdk.Coins
	Buyer sdk.AccAddress
}

//NewMsgBuyName is the constructor function for MsgBuyName
func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyName{
	return MsgBuyName{
		Name: name,
		Bid: bid,
		Buyer: buyer,

	}
}

//returns name of the module
func (msg MsgBuyName) Route() string {return "nameservice"}

//returns the action
func (msg MsgBuyName) Type() string {return "buy_name"}



//ValidateBasic runs stateless checks on the message
func (msg MsgBuyName) ValidateBasic() sdk.Error{
	if msg.Buyer.Empty(){
		return sdk.ErrUnknownRequest(msg.Buyer.String())
	}
	if len(msg.Name)==0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if !msg.Bid.IsAllPositive(){
		return sdk.ErrInsufficientCoins("Bids must  be positive")
	}
	return nil
}

//GetSignBytes encodes the message for signing
func (msg MsgBuyName) GetSignBytes() []byte {
	b, err:=json.Marshal(msg)
	if err!=nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}


//GetSigners defines whose signature is required
func(msg MsgBuyName) GetSigners() []sdk.AccAddress{
	return []sdk.AccAddress{msg.Buyer}
}
















