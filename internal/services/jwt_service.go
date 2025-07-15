package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/sunsetsavorer/eat-mate-api/internal/usecases"
)

const (
	HEADER_INDEX = iota
	PAYLOAD_INDEX
	SIGNATURE_INDEX
)

type JWTService struct {
	secret   string
	lifetime time.Duration
}

func NewJWTService(
	secret string,
	lifetime time.Duration,
) *JWTService {

	return &JWTService{
		secret,
		lifetime,
	}
}

func (s JWTService) GenerateTokenByUserID(userID int64) (usecases.Token, error) {

	header := usecases.TokenHeader{
		Alg: "HS256",
		Typ: "JWT",
	}

	jsonHeader, err := json.Marshal(&header)
	if err != nil {
		return usecases.Token{}, fmt.Errorf("an error occured while marshaling token header: %v", err)
	}

	payload := usecases.TokenPayload{
		UserID: userID,
		Exp:    time.Now().Add(s.lifetime).Unix(),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return usecases.Token{}, fmt.Errorf("an error occured while marshaling token payload: %v", err)
	}

	encodedHeader := base64URLEncode(jsonHeader)
	encodedPayload := base64URLEncode(jsonPayload)

	message := encodedHeader + "." + encodedPayload

	signature := getSignature(s.secret, message)

	encodedSignature := base64URLEncode(signature)

	token := message + "." + encodedSignature

	return usecases.Token{
		Payload: payload,
		Value:   token,
	}, nil
}

func (s JWTService) ParseToken(token string) (usecases.Token, error) {

	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		return usecases.Token{}, fmt.Errorf("invalid token format")
	}

	header := parts[HEADER_INDEX]
	payload := parts[PAYLOAD_INDEX]
	exceptedSignature := parts[SIGNATURE_INDEX]

	message := header + "." + payload

	actualSignature := getSignature(s.secret, message)
	encodedSignature := base64URLEncode(actualSignature)

	if encodedSignature != exceptedSignature {
		return usecases.Token{}, fmt.Errorf("invalid token data")
	}

	decodedPayload, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return usecases.Token{}, fmt.Errorf("failed to decode payload: %v", err)
	}

	payloadStruct := usecases.TokenPayload{}

	err = json.Unmarshal(decodedPayload, &payloadStruct)
	if err != nil {
		return usecases.Token{}, fmt.Errorf("failed to parse payload: %v", err)
	}

	if isTokenExpired(payloadStruct.Exp) {
		return usecases.Token{}, fmt.Errorf("token expired")
	}

	return usecases.Token{
		Payload: payloadStruct,
		Value:   token,
	}, nil
}

func base64URLEncode(data []byte) string {

	return base64.RawURLEncoding.EncodeToString(data)
}

func getSignature(secret, message string) []byte {

	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(message))

	signature := hash.Sum(nil)

	return signature
}

func isTokenExpired(unixTimestamp int64) bool {

	return time.Now().Unix() > unixTimestamp
}
