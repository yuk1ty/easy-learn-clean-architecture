package vo

type UserId struct {
	Value string
}

func NewUserId() *UserId {
	return &UserId{Value: "uuid"} // TODO
}