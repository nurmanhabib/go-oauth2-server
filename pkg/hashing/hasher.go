package hashing

import "context"

// Hasher is the hashing process interface.
type Hasher interface {
	Check(ctx context.Context, plain string, hashed string) bool
}
