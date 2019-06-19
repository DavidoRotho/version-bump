package dummy


// Dummy is the struct for Google cloud storage
type Dummy struct{
}

// Lock placeholder.
func (d Dummy) Lock() string {
	return "lock Dummy"
}

// Unlock google cloud storage.
func (d Dummy) Unlock() string {
	return "unlock dummy"
}
