package app


import(
	"github.com/tendermint/tendermint/libs/log"
	"github.com/cosmos/cosmos-sdk/x/auth"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	dbm "github.com/tendermint/tendermint/libs/db"

)

const(
	appName="nameservice"
)

//this type embeds baseapp, so it has access to ALL of baseapp's methods
type nameServiceApp struct{
	*bam.BaseApp
}

func NewNameServiceApp(logger log.Logger, db dbm.DB) *nameServiceApp {
	//first define the top level Codec that will be shared by many different modules, will be explained later
	cdc:= MadeCodec()

	//BaseApp will handle the interactions between Tendermint through the ABCI protocol
	bApp:=bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecorder(cdc))

	var app=&nameServiceApp{
		BaseApp: bApp,
		cdc: cdc,

	}

	return app

}




