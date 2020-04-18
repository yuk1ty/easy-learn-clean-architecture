package vo

import "github.com/google/uuid"

type RoleId struct {
	Value string
}

func NewRoleId() *RoleId {
	return &RoleId{Value: uuid.New().String()}
}
