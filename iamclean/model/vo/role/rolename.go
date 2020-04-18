package vo

type RoleName struct {
	Value string
}

func NewRoleName(name string) (*RoleName, error) {
	if len(name) > 15 {
		return nil, nil // TODO return error
	}
	return &RoleName{Value: name}, nil
}
