package main

import (
	"testing"

	"github.com/aneesh-jose/simple-server/controllers"
)

func TestAuthenticate(t *testing.T) {
	authTest := []struct {
		token    string
		username string
	}{
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFuZWVzaGpvc2UiLCJwYXNzd29yZCI6InNhbXBsZXBhc3MiLCJleHAiOjE2Mjg2NjU1Mjd9.LU90k2_qfJLINzbZbS5VkQre-5AIWlS6O2s6RAhOx4k", "aneeshjose"},
		{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1C2VybmFtZSI6ImFuZWVzaGpvc2UiLCJwYXNzd29yZCI6InNhbXBsZXBhc3MiLCJleHAiOjE2Mjg2NjU1Mjd9.LU90k2_qfJLINzbZbS5VkQre-5AIWlS6O2s6RAhOx4k", ""},
	}

	for iter, data := range authTest {
		output, _ := controllers.JWTAuthenticate(&data.token)
		if output != data.username {
			t.Errorf("Status unexpected-> expected:%v, output %v on iteration: %v", data.username, output, iter)
		}
	}
}
