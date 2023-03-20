package main

import (
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	payload := Payload{
		Exp: 1679356799,
	}
	var expected = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzkzNTY3OTl9.jEGdt7osiNFBXZ9Mbhd_MT3OOilVnzzLPvW1awhl2cY"

	// Generate the JWT
	jwt, err := GenerateJwt(payload, "super-secret")

	if err != nil {
		t.Error(err)
	}
	// Verify that the JWT is not empty
	if jwt == "" {
		t.Errorf("JWT is empty")
	}
	if jwt != expected {
		t.Errorf("Wrong jwt, got other value, then expected. Expected %s got %s", expected, jwt)
	}
}
