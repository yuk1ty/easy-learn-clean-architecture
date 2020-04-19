package apperr

type MySqlConnectionErr struct {
	msg string
}

func NewMySqlConnectionErr(msg string) *MySqlConnectionErr {
	return &MySqlConnectionErr{msg: msg}
}

func (e *MySqlConnectionErr) Error() string {
	return e.msg
}
