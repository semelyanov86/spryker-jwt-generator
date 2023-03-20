package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Payload struct {
	Exp int64 `json:"exp"`
}

type EncodingHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func GenerateJwt(payload Payload, secret string) (string, error) {
	header := EncodingHeader{
		Alg: "HS256",
		Typ: "JWT",
	}
	// Marshal the payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	headerJson, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	// Encode the payload to base64
	payloadBase64 := base64.RawURLEncoding.EncodeToString(payloadJSON)
	headerBase64 := base64.RawURLEncoding.EncodeToString(headerJson)

	// Generate the signature using HMAC SHA256
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(headerBase64 + "." + payloadBase64))
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	// Combine the payload and signature to form the JWT
	jwt := headerBase64 + "." + payloadBase64 + "." + signature

	// Print the generated JWT
	return jwt, nil
}

func main() {
	now := time.Now()
	currentYear, currentMonth, currentDay := now.Date()

	payload := Payload{
		Exp: time.Date(currentYear, currentMonth, currentDay, 22, 0, 0, 0, time.UTC).Unix(),
	}
	result, err := GenerateJwt(payload, os.Getenv("SPRYKER_JWT_SECRET"))
	if err != nil {
		fmt.Printf("Error in generating token: %s", err)
	}
	fmt.Println(result)
}
