package exceptions

type (
	ValidationError struct {
		Err []ValidationField
	}

	ValidationField struct {
		Name string
		Err  error
	}
)

func NewValidationError(err []ValidationField) *ValidationError {
	return &ValidationError{err}
}

func (e ValidationError) Error() string {
	return "validation error"
}

func (e ValidationError) Errors() map[string]string {

	errResp := make(map[string]string)

	for _, err := range e.Err {
		errResp[err.Name] = err.Err.Error()
	}

	return errResp
}
