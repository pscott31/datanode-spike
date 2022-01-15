package sqlstore

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"vega.spike/entities"
)

const connStr string = "postgres://vega:vega@localhost/vega"

type Store struct {
	db *gorm.DB
}

func NewStore() Store {
	db, err := gorm.Open(postgres.Open(connStr))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("connected")
	db.AutoMigrate(&entities.Transfer{})
	db.AutoMigrate(&entities.Account{})
	db.AutoMigrate(&entities.Asset{})
	db.AutoMigrate(&entities.Party{})
	return Store{db: db}
}

func (s Store) GetDB() *gorm.DB {
	return s.db
}
