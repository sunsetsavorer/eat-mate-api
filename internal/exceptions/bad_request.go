package exceptions

type (
	BadRequestError struct {
		Err error
	}
)

func NewBadRequestError(err error) *BadRequestError {

	return &BadRequestError{err}
}

func (e BadRequestError) Error() string {
	return e.Err.Error()
}
