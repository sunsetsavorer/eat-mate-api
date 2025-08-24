package exceptions

type (
	ForbiddenError struct {
		Err error
	}
)

func NewForbiddenError(err error) *ForbiddenError {

	return &ForbiddenError{err}
}

func (e ForbiddenError) Error() string {
	return e.Err.Error()
}
