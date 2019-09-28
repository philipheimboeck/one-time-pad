package domain

// Value is a key value pair, stored in the database
type Value struct {
	Value        []byte
	TTL          int
	Accesses     int
	HashedSecret []byte
	EncryptedKey []byte
}
