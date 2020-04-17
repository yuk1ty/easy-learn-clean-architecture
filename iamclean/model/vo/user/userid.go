package user

type Id struct {
	Value string
}

func NewId() *Id {
	return &Id{Value:"uuid"} // TODO
}