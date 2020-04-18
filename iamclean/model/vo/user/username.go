package vo

type UserName struct {
	Value string
}

func NewUserName(name string) (*UserName, error) {
	if len(name) > 15 {
		return nil, nil // TODO return error
	}
	return &UserName{Value: name}, nil
}
