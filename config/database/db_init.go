package database

import (
	"fmt"
	"go.etcd.io/bbolt"
	"log"
)

type Initializer struct {
	db *bbolt.DB
}

func NewInitializer(dbManager *DBManager) *Initializer {
	return &Initializer{db: dbManager.DB}
}

func (i *Initializer) Init() {
	i.initializeSchema()
}

func (i *Initializer) initializeSchema() {
	err := i.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("chatroom"))
		if err != nil {
			return fmt.Errorf("create chatroom bucket: %s", err)
		}

		_, err = tx.CreateBucketIfNotExists([]byte("user"))
		if err != nil {
			return fmt.Errorf("create user bucket: %s", err)
		}

		return nil
	})

	if err != nil {
		log.Printf("Failed to initialize schema: %v", err)
		panic(err)
	}
}
