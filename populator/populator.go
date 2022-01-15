package populator

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"vega.spike/entities"
	"vega.spike/sqlstore"
	"vega.spike/utils"
)

type assetGetter interface {
	GetAll() []*entities.Asset
}

type assetInserter interface {
	Insert(*entities.Asset)
}

type partyInserter interface {
	Insert(*entities.Party)
}

type accountGetter interface {
	GetAll() []*entities.Account
}

type accountInserter interface {
	Insert(*entities.Account)
}

type transferInserter interface {
	Insert(*entities.Transfer)
}

func CreateRandomParties(
	numAccounts int,
	assetStore assetGetter,
	partyStore partyInserter,
	accountStore accountInserter) {
	defer utils.TimeTrack(time.Now(), "create parties")

	assets := assetStore.GetAll()

	for i := 0; i < numAccounts; i++ {
		randomAsset := *assets[rand.Intn(len(assets))]

		randomParty := createRandomParty()
		partyStore.Insert(&randomParty)

		randomAccount := createRandomAccount(randomAsset, randomParty)
		accountStore.Insert(&randomAccount)
	}
}

func CreateRandomTransfers(
	numTransfers int,
	//transferStore transferInserter,
	transferStore *sqlstore.TransferStore,
	accountStore accountGetter,
) {
	accounts := accountStore.GetAll()

	defer utils.TimeTrack(time.Now(), fmt.Sprintf("create %d transfers", numTransfers))

	transferStore.StartTx()
	for i := 0; i < numTransfers; i++ {
		createRandomTansfer(accounts, transferStore, accountStore)
	}
	transferStore.CommitTx()

}

func createRandomTansfer(
	accounts []*entities.Account,
	transferStore transferInserter,
	accountStore accountGetter,
) {
	accountFrom := accounts[rand.Intn(len(accounts))]

	// Get a different account for the same asset
	var accountTo *entities.Account
	for accountTo == nil || accountTo.AssetID != accountFrom.AssetID {
		accountTo = accounts[rand.Intn(len(accounts))]
	}

	transfer := entities.Transfer{
		ID:            generateID(),
		AccountFromID: accountFrom.ID,
		AccountToID:   accountTo.ID,
		Quantity:      rand.Intn(100),
	}
	transferStore.Insert(&transfer)
}

func createAssets(store assetInserter) {
	nums := []string{"1", "2", "3", "4"}
	for _, num := range nums {
		store.Insert(&entities.Asset{
			ID:          "asset" + num + "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Name:        "Asset " + num,
			Symbol:      "ass" + num,
			TotalSupply: "1000000",
			Decimals:    1000000,
		})
	}
}

func generateID() string {
	currentTime := time.Now().UnixNano()
	currentTimeString := strconv.FormatInt(currentTime, 10)
	hash := sha256.Sum256([]byte(currentTimeString))
	return fmt.Sprintf("%x", hash)
}

func createRandomParty() entities.Party {
	return entities.Party{
		ID: generateID(),
	}
}

func createRandomAccount(asset entities.Asset, party entities.Party) entities.Account {
	return entities.Account{
		ID:      generateID(),
		PartyID: party.ID,
		AssetID: asset.ID,
	}
}
