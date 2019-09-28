package model

import "../domain"

// DefaultModel is the standard implementation
type DefaultModel struct {
}

// Store a new key-value pair, including the ttl and number of possible accesses
func (m DefaultModel) Store(Value domain.Value) {

}

// Get by a key and decrease the possible accesses
func (m DefaultModel) Get(key string) domain.Value {
	v := domain.Value{Key: key, Value: "value", TTL: 100, Accesses: 1}
	return v
}

// Delete a key from the database
func (m DefaultModel) Delete(key string) {

}
