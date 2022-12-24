package util

import "testing"

func TestHashPassword(t *testing.T) {
	password := "12345"
	for i := 0; i < 5; i++ {
		hash, err := HashPassword(password)
		if err != nil {
			t.Error(err)
			return
		}
		bol := CheckPasswordHash(password, hash)
		t.Logf("\n%s %+v", hash, bol)
	}
}
