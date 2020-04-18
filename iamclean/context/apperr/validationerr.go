package apperr

type ValidationErr struct {
	msg string
}

func NewValidationErr(msg string) *ValidationErr {
	return &ValidationErr{msg: msg}
}

func (e *ValidationErr) Error() string {
	return e.msg
}
