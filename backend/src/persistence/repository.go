package persistence

import "../domain"

// Repository fetches and stores values
type Repository interface {
	Get(key string) (domain.Value, error)
	Put(key string, value domain.Value)
	Delete(key string)
}
