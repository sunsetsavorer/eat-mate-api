package exceptions

type ManyRequestsError struct {
	Err error
}

func NewManyRequestsError(err error) *ManyRequestsError {

	return &ManyRequestsError{err}
}

func (e ManyRequestsError) Error() string {
	return e.Err.Error()
}
