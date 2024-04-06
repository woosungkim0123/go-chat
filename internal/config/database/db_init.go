package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/bbolt"
	"log"
	"ws/internal/auth/domain"
)

type Initializer struct {
	db *bbolt.DB
}

func NewInitializer(dbManager *DBManager) *Initializer {
	return &Initializer{db: dbManager.DB}
}

func (i *Initializer) Init() {
	i.initializeSchema()
	i.initializeData()
}

func (i *Initializer) initializeSchema() {
	err := i.db.Update(func(tx *bbolt.Tx) error {

		err := tx.DeleteBucket([]byte("chatroom"))
		if err != nil && !errors.Is(err, bbolt.ErrBucketNotFound) {
			return fmt.Errorf("delete chatroom bucket: %s", err)
		}

		err = tx.DeleteBucket([]byte("user"))
		if err != nil && !errors.Is(err, bbolt.ErrBucketNotFound) {
			return fmt.Errorf("delete user bucket: %s", err)
		}

		err = tx.DeleteBucket([]byte("chatroomMessage"))
		if err != nil && !errors.Is(err, bbolt.ErrBucketNotFound) {
			return fmt.Errorf("delete chatroomMessage bucket: %s", err)
		}

		_, err = tx.CreateBucketIfNotExists([]byte("chatroom"))
		if err != nil {
			return fmt.Errorf("create chatroom bucket: %s", err)
		}

		_, err = tx.CreateBucketIfNotExists([]byte("user"))
		if err != nil {
			return fmt.Errorf("create user bucket: %s", err)
		}

		_, err = tx.CreateBucketIfNotExists([]byte("chatroomMessage"))
		if err != nil {
			return fmt.Errorf("create chatroomMessage bucket: %s", err)
		}

		return nil
	})

	if err != nil {
		log.Printf("Failed to initialize schema: %v", err)
		panic(err)
	}
}

func (i *Initializer) initializeData() {
	users := []domain.User{
		{Name: "홍길동", LoginID: "test1", ProfileImage: "path/to/alice.jpg"},
		{Name: "김철수", LoginID: "test2", ProfileImage: "path/to/bob.jpg"},
		{Name: "이영희", LoginID: "test3", ProfileImage: "path/to/bob.jpg"},
	}

	err := i.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("user"))
		if bucket == nil {
			return fmt.Errorf("user bucket not found")
		}

		id, err := bucket.NextSequence()
		if err != nil {
			return fmt.Errorf("failed to get next user ID: %s", err)
		}

		for _, user := range users {
			user.ID = int(id)
			id++

			userData, err := json.Marshal(user)
			if err != nil {
				return fmt.Errorf("failed to marshal user data: %s", err)
			}

			err = bucket.Put([]byte(fmt.Sprintf("%d", user.ID)), userData)
			if err != nil {
				return fmt.Errorf("failed to save user data: %s", err)
			}
		}

		return nil
	})

	if err != nil {
		log.Printf("Failed to initialize user data: %v", err)
		panic(err)
	}
}
