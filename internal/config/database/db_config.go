package database

import (
	"go.etcd.io/bbolt"
	"log"
	"time"
)

type DBManager struct {
	DB *bbolt.DB
}

func NewDBManager(dbPath string) *DBManager {
	db, err := bbolt.Open(dbPath, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return &DBManager{DB: db}
}

func (manager *DBManager) Close() {
	err := manager.DB.Close()
	if err != nil {
		log.Fatal("Failed to close the database:", err)
	}
}
