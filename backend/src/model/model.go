package model

import "../domain"

// Model to store, get and delete values
type Model interface {
	Store(Value domain.Value)
	Get(key string) domain.Value
	Delete(key string)
}
