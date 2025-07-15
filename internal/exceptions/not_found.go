package exceptions

type NotFoundError struct {
	Err error
}

func NewNotFoundError(err error) *NotFoundError {

	return &NotFoundError{err}
}

func (e NotFoundError) Error() string {
	return e.Err.Error()
}
