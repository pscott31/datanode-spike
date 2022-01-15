package sqlstore

import (
	"vega.spike/entities"
)

type AssetStore struct {
	Store
}

func NewAssetStore(base Store) *AssetStore {
	store := AssetStore{
		Store: base,
	}
	return &store
}

func (as *AssetStore) Insert(a *entities.Asset) {
	as.db.Create(&a)
}

func (as *AssetStore) GetAll() []*entities.Asset {
	var assets []*entities.Asset
	as.db.Find(&assets)
	return assets
}
