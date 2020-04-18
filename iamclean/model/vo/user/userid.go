package vo

import "github.com/google/uuid"

type UserId struct {
	Value string
}

func NewUserId() *UserId {
	return &UserId{Value: uuid.New().String()}
}
