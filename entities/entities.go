package entities

type Asset struct {
	ID                     string `gorm:"primarykey"`
	Name                   string `gorm:"unique"`
	Symbol                 string
	TotalSupply            string `gorm:"type:decimal(30,0)"` // todo: big? decimal?
	Decimals               uint64
	MinLpStake             uint64
	Source                 string // TODO: Enum? Make it polymorphic?
	Erc20Contract          string
	BuiltinMaxFaucetAmount uint64

	// TODO - method to convert to/from asset type?
}

type Party struct {
	ID string `gorm:"primarykey"`
}

type Account struct {
	ID      string `gorm:"primarykey"`
	PartyID string
	AssetID string
	Party   Party
	Asset   Asset
}

type Transfer struct {
	ID            string `gorm:"primarykey"`
	AccountFromID string
	AccountToID   string
	Quantity      int
	AccountFrom   Account
	AccountTo     Account
}
