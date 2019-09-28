package model

import (
	"errors"

	"../domain"
	"../dto"
	"../encryption"
	"../persistence"
)

// DefaultModel is the standard implementation
type DefaultModel struct {
	repository persistence.Repository
}

// MakeDefaultModel creates a new instance
func MakeDefaultModel(r persistence.Repository) DefaultModel {
	return DefaultModel{repository: r}
}

// Store a new key-value pair, including the ttl and number of possible accesses
func (m *DefaultModel) Store(key string, secret string, value dto.ValueDTO) {

	entity := domain.Value{TTL: value.TTL, Accesses: value.Accesses}

	encrypt(secret, value.Value, &entity)

	m.repository.Put(key, entity)
}

// Get by a key and decrease the possible accesses
func (m *DefaultModel) Get(key string, secret string) (dto.ValueDTO, error) {
	var value dto.ValueDTO

	entity, err := m.repository.Get(key)
	if err != nil {
		return value, err
	}

	value.Value = decrypt(secret, &entity)
	value.Accesses = entity.Accesses
	value.TTL = entity.TTL

	entity.Accesses--
	if entity.Accesses <= 0 {
		m.repository.Delete(key)
	}

	if entity.Accesses < 0 {
		return value, errors.New("Value already expired")
	}

	return value, nil
}

// Delete a key from the database
func (m *DefaultModel) Delete(key string) {
	m.repository.Delete(key)
}

func encrypt(secret string, plaintext string, value *domain.Value) {

	ciphertext, encryptedKey, hashedSecret := encryption.Encrypt(secret, plaintext)

	value.HashedSecret = hashedSecret
	value.EncryptedKey = encryptedKey
	value.Value = ciphertext
}

func decrypt(secret string, value *domain.Value) string {
	return encryption.Decrypt(secret, value.Value, value.HashedSecret, value.EncryptedKey)
}
