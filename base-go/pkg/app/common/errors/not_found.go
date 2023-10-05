package errors

const NotFoundErrorDefaultMessage = "not_found"

type NotFoundError struct {
	msg string
}

func NewNotFoundError(msg string) *NotFoundError {
	if msg == "" {
		msg = NotFoundErrorDefaultMessage
	}
	return &NotFoundError{msg: msg}
}

func (e *NotFoundError) Error() string {
	return e.msg
}
