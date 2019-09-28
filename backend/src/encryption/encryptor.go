package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// GenerateSecretKey generates a random secret to be used with AES256
func GenerateSecretKey() []byte {
	token := make([]byte, 32)
	rand.Read(token)

	return token
}

// Encrypt with AES256
func Encrypt(secret string, plaintext string) ([]byte, []byte, []byte, error) {
	secretKey := GenerateSecretKey()
	ciphertext, error := encryptString(secretKey, []byte(plaintext))
	if error != nil {
		return nil, nil, nil, error
	}

	sha256Secret := sha256.Sum256([]byte(secret))
	encryptedKey, error := encryptString(sha256Secret[0:], secretKey)
	if error != nil {
		return nil, nil, nil, error
	}

	hashedSecret, error := bcrypt.GenerateFromPassword([]byte(sha256Secret[0:]), bcrypt.MinCost)
	if error != nil {
		return nil, nil, nil, error
	}

	return ciphertext, encryptedKey, hashedSecret, nil
}

func encryptString(key []byte, plaintext []byte) ([]byte, error) {
	cipherBlock, error := getCipher(key)
	if error != nil {
		return nil, error
	}

	gcm, error := cipher.NewGCM(cipherBlock)
	if error != nil {
		return nil, error
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, error := io.ReadFull(rand.Reader, nonce); error != nil {
		return nil, error
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// Decrypt with AES256
func Decrypt(secret string, encryptedText []byte, hashedSecret []byte, encryptedKey []byte) (string, error) {

	sha256Secret := sha256.Sum256([]byte(secret))
	compareSecrets(sha256Secret[0:], []byte(hashedSecret))

	secretKey, err := decryptString(sha256Secret[0:], []byte(encryptedKey))
	if err != nil {
		return "", err
	}
	plainText, err := decryptString(secretKey, []byte(encryptedText))
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func decryptString(key []byte, ciphertext []byte) ([]byte, error) {
	cipherBlock, err := getCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("Ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func getCipher(key []byte) (cipher.Block, error) {

	cipher, error := aes.NewCipher(key)
	if error != nil {
		return nil, error
	}
	return cipher, nil
}

func compareSecrets(secret []byte, hashedSecret []byte) error {
	error := bcrypt.CompareHashAndPassword(hashedSecret, secret)
	if error != nil {
		return error
	}
	return nil
}
