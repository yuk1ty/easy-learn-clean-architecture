package role

type Name struct {
	Name string
}

func ValidateWith(name string) (*Name, error) {
	if len(name) > 15 {
		return nil, nil // TODO return error
	}
	return &Name{Name: name}, nil
}
