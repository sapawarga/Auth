package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

var tokenTest = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE2MjQ1MDY2MzAsImlzcyI6Imh0dHA6XC9cL3NhcGF3YXJnYS1zdGFnaW5nLmphYmFycHJvdi5nby5pZCIsImF1ZCI6Imh0dHA6XC9cL3NhcGF3YXJnYS1zdGFnaW5nLmphYmFycHJvdi5nby5pZCIsIm5iZiI6MTYyNDUwNjYzMCwiZXhwIjoxNjI3MDk4NjMwLCJkYXRhIjp7InVzZXJuYW1lIjoicm9ib3Rwcm92LmpkcyIsInJvbGVMYWJlbCI6IlN0YWYgUHJvdmluc2kiLCJsYXN0TG9naW5BdCI6eyJleHByZXNzaW9uIjoiVU5JWF9USU1FU1RBTVAoKSIsInBhcmFtcyI6W119fSwianRpIjoxODYwNjN9._kw2xsTL6QgUVLcbknqkNfFgVagIL8nlP3008FpztwA"

func TestParsingToken(t *testing.T) {
	parsedToken, err := jwt.Parse(tokenTest, func(t *jwt.Token) (interface{}, error) {
		return []byte("someSecretKey"), nil
	})

	if err != nil {
		t.Error(err)
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		t.Error(errors.New("token_not_valid"))
	}

	fmt.Println(claims["data"])
	fmt.Println("hello")
}
