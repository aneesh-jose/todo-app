package main

import (
	"testing"

	authentication "github.com/aneesh-jose/simple-server/utils/auth"
)

func TestAuthenticate(t *testing.T) {
	authTest := []struct {
		testType string
		number   int
		token    string
		username string
	}{
		{"OK", 1, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFuZWVzaGpvc2UiLCJwYXNzd29yZCI6InNhbXBsZXBhc3MiLCJleHAiOjE2Mjg2NjU1Mjd9.LU90k2_qfJLINzbZbS5VkQre-5AIWlS6O2s6RAhOx4k", "aneeshjose"},
		{"ERROR", 2, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1C2VybmFtZSI6ImFuZWVzaGpvc2UiLCJwYXNzd29yZCI6InNhbXBsZXBhc3MiLCJleHAiOjE2Mjg2NjU1Mjd9.LU90k2_qfJLINzbZbS5VkQre-5AIWlS6O2s6RAhOx4k", ""},
	}

	for iter, data := range authTest {
		output, err := authentication.JWTAuthenticate(&data.token)
		if err != nil {
			if data.testType != "ERROR" {
				t.Errorf("Error occured in %v as %v", data.number, err)
			}
		}
		if output != data.username {
			t.Errorf("Status unexpected-> expected:%v, output %v on iteration: %v", data.username, output, iter)
		}
	}
}
