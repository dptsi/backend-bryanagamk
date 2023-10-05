package errors

type InvariantError struct {
	msg string
}

func NewInvariantError(msg string) *InvariantError {
	return &InvariantError{msg: msg}
}

func (e *InvariantError) Error() string {
	return e.msg
}
