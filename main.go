package main

import (
	"log"
	"time"

	"vega.spike/populator"
	"vega.spike/sqlstore"
	"vega.spike/utils"
)

func main() {

	store := sqlstore.NewStore()
	assetStore := sqlstore.NewAssetStore(store)
	accountStore := sqlstore.NewAccountStore(store)
	partyStore := sqlstore.NewPartyStore(store)
	transferStore := sqlstore.NewTransferStore(store)

	defer utils.TimeTrack(time.Now(), "total time")
	//populator.createAssets(as)
	//populator.CreateRandomParties(10000, assetStore, partyStore, accountStore)
	populator.CreateRandomTransfers(10000, transferStore, accountStore)

	log.Println("done")

	_ = assetStore
	_ = partyStore
}
