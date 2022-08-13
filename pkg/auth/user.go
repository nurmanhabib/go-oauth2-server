package auth

// User is an interface for user auth.
type User interface {
	GetAuthIdentifier() string
	GetAuthPassword() string
}
