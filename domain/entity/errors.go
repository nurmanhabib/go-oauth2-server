package entity

// Error is new type for error.
type Error string

// Error is a function that returns error details.
func (e Error) Error() string {
	return string(e)
}
