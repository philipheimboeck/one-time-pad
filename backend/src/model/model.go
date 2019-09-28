package model

import "../dto"

// Model to store, get and delete values
type Model interface {
	Store(key string, secret string, value dto.ValueDTO) error
	Get(key string, secret string) (dto.ValueDTO, error)
	Delete(key string) error
}
