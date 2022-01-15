package sqlstore

import (
	"gorm.io/gorm"
	"vega.spike/entities"
)

type TransferStore struct {
	Store
	tx *gorm.DB
}

func NewTransferStore(base Store) *TransferStore {
	store := TransferStore{
		Store: base,
	}
	return &store
}

func (ts *TransferStore) Insert(t *entities.Transfer) {
	//ts.db.Create(&t)
	if err := ts.tx.Create(&t).Error; err != nil {
		panic(err)
	}
}

func (ts *TransferStore) GetAll() []*entities.Transfer {
	var transfers []*entities.Transfer
	ts.db.Find(&transfers)
	return transfers
}

// Testing
func (ts *TransferStore) StartTx() {
	ts.tx = ts.db.Begin()
}

func (ts *TransferStore) CommitTx() {
	if err := ts.tx.Commit().Error; err != nil {
		panic(err)
	}
}
