package auth

import (
	"context"
	"errors"
)

// User represents a user principal.
type User struct {
	ID    string
	Roles []string
}

// CtxKey is the key for the user in the context.
var CtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// ValidateJWT is a placeholder for JWT validation.
func ValidateJWT(ctx context.Context, token string) (context.Context, error) {
	if token == "" {
		return ctx, errors.New("token is empty")
	}

	// In a real implementation, you would validate the token
	// and extract the user information.
	user := &User{
		ID:    "user-123",
		Roles: []string{"user"},
	}

	return context.WithValue(ctx, CtxKey, user), nil
}
