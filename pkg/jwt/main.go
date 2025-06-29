package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	HEADER_INDEX = iota
	PAYLOAD_INDEX
	SIGNATURE_INDEX
)

func GenerateToken(userID int64, secret string, lifeTime time.Duration) (Token, error) {

	header := TokenHeader{
		Alg: "HS256",
		Typ: "JWT",
	}

	jsonHeader, err := json.Marshal(&header)
	if err != nil {
		return Token{}, fmt.Errorf("an error occured while marshaling token header: %v", err)
	}

	payload := TokenPayload{
		UserID: userID,
		Exp:    time.Now().Add(lifeTime).Unix(),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return Token{}, fmt.Errorf("an error occured while marshaling token payload: %v", err)
	}

	encodedHeader := base64URLEncode(jsonHeader)
	encodedPayload := base64URLEncode(jsonPayload)

	message := encodedHeader + "." + encodedPayload

	signature := getSignature(secret, message)

	encodedSignature := base64URLEncode(signature)

	token := message + "." + encodedSignature

	return Token{
		Payload: payload,
		Value:   token,
	}, nil
}

func ParseToken(token, secret string) (Token, error) {

	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		return Token{}, fmt.Errorf("invalid token format")
	}

	header := parts[HEADER_INDEX]
	payload := parts[PAYLOAD_INDEX]
	exceptedSignature := parts[SIGNATURE_INDEX]

	message := header + "." + payload

	actualSignature := getSignature(secret, message)
	encodedSignature := base64URLEncode(actualSignature)

	if encodedSignature != exceptedSignature {
		return Token{}, fmt.Errorf("invalid token data")
	}

	decodedPayload, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return Token{}, fmt.Errorf("failed to decode payload: %v", err)
	}

	payloadStruct := TokenPayload{}

	err = json.Unmarshal(decodedPayload, &payloadStruct)
	if err != nil {
		return Token{}, fmt.Errorf("failed to parse payload: %v", err)
	}

	if isTokenExpired(payloadStruct.Exp) {
		return Token{}, fmt.Errorf("token expired")
	}

	return Token{
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
