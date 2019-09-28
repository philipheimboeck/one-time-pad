package dto

// ValueDTO is needed for communication with the API
type ValueDTO struct {
	Value    string
	TTL      int
	Accesses int
}
