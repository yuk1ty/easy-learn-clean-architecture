package user

type Name struct {
	Value string
}

func ValidateWith(name string) (*Name, error) {
	if len(name) > 15 {
		return nil, nil // TODO return error
	}
	return &Name{Value: name}, nil
}
