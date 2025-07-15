package exceptions

type RepositoryError struct {
	Err error
}

func NewRepositoryError(err error) *RepositoryError {

	return &RepositoryError{err}
}

func (e RepositoryError) Error() string {
	return e.Err.Error()
}
