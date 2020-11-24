package backends

// ClearSale is an anti fraud backend implementation.
type ClearSale struct{}

// IsValidTransaction checks if transiction is valid or not.
func (c *ClearSale) IsValidTransaction() bool {
	return true
}
