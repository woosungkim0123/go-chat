package chatroom

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"log"
)

type Repository struct {
	db              *bbolt.DB
	chatroom        string
	chatroomMessage string
}

func NewRepository(db *bbolt.DB) *Repository {
	return &Repository{db: db, chatroom: "chatroom", chatroomMessage: "chatroom_message"}
}

func (r *Repository) GetChatroomListByUserId(userID int) []ChatroomWithLastMessageDTO {
	var chatroomListDto []ChatroomWithLastMessageDTO

	err := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(r.chatroom))
		if b == nil {
			return fmt.Errorf("bucket not found: %s", r.chatroom)
		}

		return b.ForEach(func(k, v []byte) error {
			var room Chatroom
			if err := json.Unmarshal(v, &room); err != nil {
				return fmt.Errorf("failed to unmarshal chatroom: %v", err)
			}

			if room.Type == Single && r.containsParticipant(room.Participants, userID) {
				lastMsg := r.getLastMessage(room.ID)
				chatroomDto := NewChatroomWithLastMessageDTO(&room, lastMsg)
				chatroomListDto = append(chatroomListDto, *chatroomDto)
			}

			return nil
		})
	})

	if err != nil {
		log.Printf("Failed to get chatroom list by user id: %v", err)
		panic(err)
	}

	return chatroomListDto
}

func (r *Repository) containsParticipant(participants []ChatroomParticipant, userID int) bool {
	for _, p := range participants {
		if p.Id == userID {
			return true
		}
	}
	return false
}

func (r *Repository) getLastMessage(roomID int) *ChatroomMessage {
	var lastMessage *ChatroomMessage

	err := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(r.chatroomMessage))
		if b == nil {
			return fmt.Errorf("bucket not found: %s", r.chatroomMessage)
		}

		return b.ForEach(func(k, v []byte) error {
			var msg ChatroomMessage
			if err := json.Unmarshal(v, &msg); err != nil {
				return err
			}
			if msg.RoomID == roomID {
				lastMessage = &msg
			}
			return nil
		})
	})

	if err != nil {
		log.Printf("Failed to get last message: %v", err)
		panic(err)
	}

	return lastMessage
}
