package sqlstore

import (
	"vega.spike/entities"
)

type AccountStore struct {
	Store
}

func NewAccountStore(base Store) *AccountStore {
	store := AccountStore{
		Store: base,
	}
	return &store
}

func (as *AccountStore) Insert(a *entities.Account) {
	as.db.Create(&a)
}

func (as *AccountStore) GetAll() []*entities.Account {
	var accounts []*entities.Account
	as.db.Find(&accounts)
	return accounts
}
