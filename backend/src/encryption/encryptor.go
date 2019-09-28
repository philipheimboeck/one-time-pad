package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
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
func Encrypt(secret string, plaintext string) ([]byte, []byte, []byte) {
	secretKey := GenerateSecretKey()
	ciphertext := encryptString(secretKey, []byte(plaintext))

	sha256Secret := sha256.Sum256([]byte(secret))
	encryptedKey := encryptString(sha256Secret[0:], secretKey)

	hashedSecret, error := bcrypt.GenerateFromPassword([]byte(sha256Secret[0:]), bcrypt.MinCost)
	if error != nil {
		panic(error)
	}

	return ciphertext, encryptedKey, hashedSecret
}

func encryptString(key []byte, plaintext []byte) []byte {
	cipherBlock := getCipher(key)

	gcm, error := cipher.NewGCM(cipherBlock)
	if error != nil {
		panic(error)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, error := io.ReadFull(rand.Reader, nonce); error != nil {
		panic(error)
	}

	return gcm.Seal(nonce, nonce, plaintext, nil)
}

// Decrypt with AES256
func Decrypt(secret string, encryptedText []byte, hashedSecret []byte, encryptedKey []byte) string {

	sha256Secret := sha256.Sum256([]byte(secret))
	compareSecrets(sha256Secret[0:], []byte(hashedSecret))

	secretKey := decryptString(sha256Secret[0:], []byte(encryptedKey))
	plainText := decryptString(secretKey, []byte(encryptedText))

	return string(plainText)
}

func decryptString(key []byte, ciphertext []byte) []byte {
	cipherBlock := getCipher(key)

	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
		panic(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return plaintext
}

func getCipher(key []byte) cipher.Block {

	cipher, error := aes.NewCipher(key)
	if error != nil {
		panic(error)
	}
	return cipher
}

func compareSecrets(secret []byte, hashedSecret []byte) {
	error := bcrypt.CompareHashAndPassword(hashedSecret, secret)
	if error != nil {
		panic(error)
	}
}
