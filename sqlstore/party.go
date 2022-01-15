package sqlstore

import (
	"vega.spike/entities"
)

type PartyStore struct {
	Store
}

func NewPartyStore(base Store) *PartyStore {
	store := PartyStore{
		Store: base,
	}
	return &store
}

func (ps *PartyStore) Insert(a *entities.Party) {
	ps.db.Create(&a)
}

func (ps *PartyStore) GetAll() []*entities.Party {
	var parties []*entities.Party
	ps.db.Find(&parties)
	return parties
}
