package password

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
)

// EncryptPassword takes a string encrypts and returns the ciphertext
func EncryptPassword(secretKey string, password string) (string, []byte, error) {
	//setup encryption for storing password
	keyCipherBlock, err := getEncryptionKey(secretKey)
	if err != nil {
		return "", []byte(""), err
	}

	//setup nonce
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", []byte(""), err
	}

	//encrypt password
	aesgcm, err := cipher.NewGCM(keyCipherBlock)
	if err != nil {
		return "", []byte(""), err
	}
	cipherPassword := aesgcm.Seal(nil, nonce, []byte(password), nil)

	return string(cipherPassword), []byte(nonce), nil
}

// DecryptPassword reads a ciphertext and returns the password
func DecryptPassword(secretKey string, nonce []byte, cipherPassword string) (string, error) {
	//setup encryption for storing password
	keyCipherBlock, err := getEncryptionKey(secretKey)
	if err != nil {
		return "", err
	}

	//decrypt password
	aesgcm, err := cipher.NewGCM(keyCipherBlock)
	if err != nil {
		return "", err
	}
	password, err := aesgcm.Open(nil, nonce, []byte(cipherPassword), nil)
	if err != nil {
		return "", err
	}

	return string(password), nil
}

func getEncryptionKey(key string) (cipher.Block, error) {
	if len(key) <= 0 {
		return nil, errors.New("SECRET_KEY environment variable not found")
	}

	hashedKey := hash(key)

	keyCipherBlock, err := aes.NewCipher(hashedKey)
	if err != nil {
		return nil, err
	}

	return keyCipherBlock, nil
}

func hash(str string) []byte {
	h := sha256.Sum256([]byte(str))
	return h[:]
}
