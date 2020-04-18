package vo

type RoleId struct {
	Value string
}

func NewRoleId() *RoleId {
	return &RoleId{Value: "uuid"} // TODO
}