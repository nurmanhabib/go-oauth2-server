package auth

// Error is an error type in auth.
type Error string

// Error is a function to retrieve error messages.
func (e Error) Error() string {
	return string(e)
}

// ErrXXX is an error message constant used in authentication.
var (
	ErrInvalidCredentials = Error("auth.invalid_credentials")
	ErrUserNotFound       = Error("auth.user_not_found")
)
