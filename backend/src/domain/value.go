package domain

// Value is a key value pair, stored in the database
type Value struct {
	Key      string
	Value    string
	TTL      int
	Accesses int
}
