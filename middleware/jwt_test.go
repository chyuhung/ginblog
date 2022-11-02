package middleware

import (
	"fmt"
	"testing"
)

func TestSetToken(t *testing.T) {
	var token string
	username := "test"
	token, errcode := SetToken(username)
	fmt.Println(errcode)
	fmt.Println(token)
}
