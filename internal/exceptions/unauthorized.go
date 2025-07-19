package exceptions

type UnauthorizedError struct {
	Err error
}

func NewUnauthorizedError(err error) *UnauthorizedError {

	return &UnauthorizedError{err}
}

func (e UnauthorizedError) Error() string {
	return e.Err.Error()
}
