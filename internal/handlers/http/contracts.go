package http

type (
	OtherError struct {
		Other string `json:"other"`
	}

	ErrorResp[T OtherError | map[string]string] struct {
		Errors T `json:"errors"`
	}

	SuccessResp struct {
		Success bool `json:"success"`
	}

	SuccessDataResp struct {
		Data any `json:"data"`
	}
)

const (
	TypeOtherError      = "other"
	TypeValidationError = "validation"
	TypeSuccessResp     = "success"
	TypeDataResp        = "data"
)
