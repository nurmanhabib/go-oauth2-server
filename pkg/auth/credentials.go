package auth

const (
	// CredentialsPasswordKey is a constant for the reserved word for the password in a credential.
	CredentialsPasswordKey = "password"
)

// Credentials are a key value pair, at least have a key password if authentication is to be used.
// As the username key is not specified, each case can have a different key. For example email as username.
type Credentials map[string]interface{}

// WithoutPassword is a function to retrieve credentials with no password. Usually used to retrieve users by username.
// In this case it is possible to query with where in credentials.
func (c Credentials) WithoutPassword() Credentials {
	credentials := Credentials{}

	for k, v := range c {
		if k == CredentialsPasswordKey {
			continue
		}

		credentials[k] = v
	}

	return credentials
}

// HasPassword is a function to ensure the key password in the credentials.
func (c Credentials) HasPassword() bool {
	_, ok := c[CredentialsPasswordKey]
	return ok
}

// GetPassword is a function to retrieve the password in the credentials.
func (c Credentials) GetPassword() string {
	if v, ok := c[CredentialsPasswordKey].(string); ok {
		return v
	}
	return ""
}
