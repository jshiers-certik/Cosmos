package nameservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//Handler determines what actions need to be taken when a message is received

//Newhandler returns a handler for "nameservice" type messages. It is essentially a
//sub-router that directs messages coming into this module to the proper handler.
//At the moment, there is only 1 Msg Handler

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

//Now, need to define logic for handling the MsgSetName message in handleMsgSetName:
func handleMsgSetName(ctx sdk.Context, keeper Keeper, msg MsgSetName) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) { //checks if the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result()
	}
	keeper.SetName(ctx, msg.Name, msg.Value) //if so, set the name to the value specified in the message
	return sdk.Result{}
}

//handle a message to buyname, performs the state transitions triggered by the message.
//Also need another ValidateBasic function that can query the application state (i.e., isn't stateless)
func handleMsgBuyName(ctx sdk.Context, keeper Keeper, msg MsgBuyName) sdk.Result {
	if keeper.GetPrice(ctx, msg.Name).IsAllGT(msg.Bid) { //checks if bid price > than price paid by current owner
		return sdk.ErrInsufficientCoins("Bid not high enough") / Result()
	}
	if keeper.HasOwner(ctx, msg.Name) {
		_, err := keeper.coinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	} else {
		_, _, err := keeper.coinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid) //If so, deduct the bid amount from the sender
		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	}
	keeper.SetOwner(ctx, msg.Name, msg.Buyer)
	keeper.SetPrice(ctx, msg.Name, msg.Bid)
	return sdk.Result{}
}
