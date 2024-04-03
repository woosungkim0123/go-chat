package chatroom

import "go.etcd.io/bbolt"

type Repository struct {
	db *bbolt.DB
}

func NewRepository(db *bbolt.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetChatroomListByUserId(userId int) []Chatroom {
	return nil
}
