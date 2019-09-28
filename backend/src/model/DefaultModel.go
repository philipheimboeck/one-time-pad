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
func (m *DefaultModel) Store(key string, secret string, value dto.ValueDTO) error {

	entity := domain.Value{TTL: value.TTL, Accesses: value.Accesses}

	err := encrypt(secret, value.Value, &entity)
	if err != nil {
		return err
	}

	return m.repository.Put(key, entity)
}

// Get by a key and decrease the possible accesses
func (m *DefaultModel) Get(key string, secret string) (dto.ValueDTO, error) {
	var value dto.ValueDTO

	entity, err := m.repository.Get(key)
	if err != nil {
		return value, err
	}

	dcryptedValue, err := decrypt(secret, &entity)
	if err != nil {
		return value, err
	}

	entity.Accesses--
	ttl, err := m.repository.GetTTL(key)
	if err != nil {
		return value, err
	}

	entity.TTL = ttl

	if entity.Accesses <= 0 || entity.TTL <= 0 {
		m.repository.Delete(key)
	} else {
		err = m.repository.Put(key, entity)
		if err != nil {
			return value, err
		}
	}

	if entity.Accesses < 0 {
		return value, errors.New("Value already expired")
	}

	value.Value = dcryptedValue
	value.Accesses = entity.Accesses
	value.TTL = entity.TTL

	return value, nil
}

// Delete a key from the database
func (m *DefaultModel) Delete(key string) error {
	return m.repository.Delete(key)
}

func encrypt(secret string, plaintext string, value *domain.Value) error {

	ciphertext, encryptedKey, hashedSecret, error := encryption.Encrypt(secret, plaintext)

	value.HashedSecret = hashedSecret
	value.EncryptedKey = encryptedKey
	value.Value = ciphertext

	return error
}

func decrypt(secret string, value *domain.Value) (string, error) {
	return encryption.Decrypt(secret, value.Value, value.HashedSecret, value.EncryptedKey)
}
