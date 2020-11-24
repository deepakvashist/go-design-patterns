package backends

// Serasa is an anti fraud backend implementation.
type Serasa struct{}

// IsValidTransaction checks if transiction is valid or not.
func (s *Serasa) IsValidTransaction() bool {
	return true
}
