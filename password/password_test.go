package password

import (
	"testing"
)

var (
	secretKey      string
	cipherPassword string
	password       string
	nonce          []byte
)

func TestEncryptPassword(t *testing.T) {
	//arrange
	secretKey = "fjeifjeifjeifjeisldkflsldkflsldl"

	//act
	c, n, err := EncryptPassword(secretKey, "ThisIsATest")
	cipherPassword = c
	nonce = n

	//assert
	if err != nil {
		t.Error("Did not return nil for error")
	}

	if len(cipherPassword) <= 0 {
		t.Error("Did not return a ciphertext")
	} else {
		t.Log(cipherPassword)
	}
}

func TestDecryptPassword(t *testing.T) {
	//arrange

	//act
	password, err := DecryptPassword(secretKey, nonce, cipherPassword)

	//assert
	if err != nil {
		t.Error("Did not return nil for error")
	}

	if password != "ThisIsATest" {
		t.Log(password)
		t.Error("Did not retrieve password")
	}
}
