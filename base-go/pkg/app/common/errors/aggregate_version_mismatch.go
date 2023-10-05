package errors

const (
	AggregateVersionMismatchErrorMessage = "aggregate_version_mismatch"
)

type AggregateVersionMismatchError struct {
	msg string
}

func NewAggregateVersionMismatchError() *AggregateVersionMismatchError {
	return &AggregateVersionMismatchError{AggregateVersionMismatchErrorMessage}
}

func (e *AggregateVersionMismatchError) Error() string {
	return e.msg
}
