package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/bbolt"
	"log"
	"ws/internal/chatroom/domain"
	"ws/internal/chatroom/dto"
	apperror2 "ws/internal/common/apperror"
	"ws/internal/common/util"
)

type ChatroomRepository struct {
	db              *bbolt.DB
	chatroom        string
	chatroomMessage string
}

func NewRepository(db *bbolt.DB) *ChatroomRepository {
	return &ChatroomRepository{db: db, chatroom: "chatroom", chatroomMessage: "chatroom_message"}
}

func (r *ChatroomRepository) GetChatroomListByUserID(userID int) []dto.ChatroomWithLastMessageDTO {
	var chatroomListDto []dto.ChatroomWithLastMessageDTO

	err := r.db.View(func(tx *bbolt.Tx) error {
		b, bucketError := util.GetBucket(tx, r.chatroom)
		if bucketError != nil {
			return bucketError
		}

		return b.ForEach(func(k, v []byte) error {
			var room domain.Chatroom
			if err := json.Unmarshal(v, &room); err != nil {
				return fmt.Errorf("failed to unmarshal chatroom: %v", err)
			}

			if (room.Type == domain.Single || room.Type == domain.Mine) && r.containsParticipant(room.Participants, userID) {
				lastMsg := r.getLastMessage(room.ID)
				chatroomDto := dto.NewChatroomWithLastMessageDTO(&room, lastMsg)
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

func (r *ChatroomRepository) GetMineChatroom(userID int) (*domain.Chatroom, *apperror2.CustomError) {
	var mineChatroom *domain.Chatroom
	var FoundChatroomAndStopIterator = errors.New("found chatroom")

	dbError := r.db.View(func(tx *bbolt.Tx) error {
		b, bucketError := util.GetBucket(tx, r.chatroom)
		if bucketError != nil {
			return bucketError
		}

		chatroomResult := b.ForEach(func(k, v []byte) error {
			var room domain.Chatroom
			if jsonError := json.Unmarshal(v, &room); jsonError != nil {
				return util.HandleError(fmt.Sprintf("failed to unmarshal chatroom data: %v", jsonError), apperror2.FailJsonUnmarshal)
			}

			if room.Type == domain.Mine && r.containsParticipant(room.Participants, userID) {
				mineChatroom = &room
				return FoundChatroomAndStopIterator
			}

			return nil
		})

		if chatroomResult == nil {
			return util.HandleError(fmt.Sprintf("failed to find mine chatroom with user id: %d", userID), apperror2.NotFoundMineChatroom)
		} else if errors.Is(chatroomResult, FoundChatroomAndStopIterator) {
			return nil // 채팅방을 찾았으므로, 에러 없음
		} else {
			return chatroomResult // 다른 에러 처리
		}
	})

	if dbError != nil {
		if dbError.Error() == string(apperror2.NotFoundMineChatroom) {
			return nil, &apperror2.CustomError{Code: apperror2.NotFoundMineChatroom, Message: "자신의 채팅방을 찾을 수 없습니다."}
		}
		return nil, &apperror2.CustomError{Code: apperror2.DataBaseProblem, Message: "데이터베이스에 문제가 발생했습니다."}
	}

	return mineChatroom, nil
}

func (r *ChatroomRepository) AddChatroom(chatroom *domain.Chatroom) *apperror2.CustomError {
	dbError := r.db.Update(func(tx *bbolt.Tx) error {
		b, bucketError := util.GetBucket(tx, r.chatroom)
		if bucketError != nil {
			return bucketError
		}

		id, sequenceError := b.NextSequence()
		if sequenceError != nil {
			log.Printf("failed to get next sequence: %v", sequenceError)
			return sequenceError
		}
		chatroom.ID = int(id)

		chatroomData, jsonError := json.Marshal(chatroom)
		if jsonError != nil {
			log.Printf("failed to marshal chatroom data: %v", jsonError)
			return jsonError
		}

		idKey := fmt.Sprintf("%d", chatroom.ID)
		if err := b.Put([]byte(idKey), chatroomData); err != nil {
			log.Printf("failed to put chatroom: %v", err)
			return err
		}

		return nil
	})

	if dbError != nil {
		log.Printf("failed to add chatroom: %v", dbError)
		return &apperror2.CustomError{Code: apperror2.DataBaseProblem, Message: "데이터베이스에 문제가 발생했습니다."}
	}

	return nil
}

// GetChatroomMessages 채팅방 ID로 메시지 가져오기
func (r *ChatroomRepository) GetChatroomMessages(chatroomID int) ([]domain.ChatroomMessage, *apperror2.CustomError) {
	var messages []domain.ChatroomMessage

	dbError := r.db.View(func(tx *bbolt.Tx) error {
		b, bucketError := util.GetBucket(tx, r.chatroom)
		if bucketError != nil {
			return bucketError
		}

		return b.ForEach(func(k, v []byte) error {
			var msg domain.ChatroomMessage
			if jsonError := json.Unmarshal(v, &msg); jsonError != nil {
				return util.HandleError(fmt.Sprintf("failed to unmarshal chatroom data: %v", jsonError), apperror2.FailJsonUnmarshal)
			}

			if msg.RoomID == chatroomID {
				messages = append(messages, msg)
			}
			return nil
		})
	})

	if dbError != nil {
		return nil, &apperror2.CustomError{Code: apperror2.DataBaseProblem, Message: "데이터베이스에서 유저를 찾는 중 문제가 발생했습니다."}
	}

	return messages, nil
}

func (r *ChatroomRepository) containsParticipant(participants []domain.ChatroomParticipant, userID int) bool {
	for _, p := range participants {
		if p.ID == userID {
			return true
		}
	}
	return false
}

func (r *ChatroomRepository) getLastMessage(roomID int) *domain.ChatroomMessage {
	var lastMessage *domain.ChatroomMessage

	err := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(r.chatroomMessage))
		if b == nil {
			return fmt.Errorf("bucket not found: %s", r.chatroomMessage)
		}

		return b.ForEach(func(k, v []byte) error {
			var msg domain.ChatroomMessage
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
