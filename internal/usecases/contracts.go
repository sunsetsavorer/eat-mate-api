package usecases

type (
	Token struct {
		Payload TokenPayload `json:"payload"`
		Value   string       `json:"value"`
	}

	TokenHeader struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}

	TokenPayload struct {
		UserID int64 `json:"user_id"`
		Exp    int64 `json:"exp"`
	}
)
