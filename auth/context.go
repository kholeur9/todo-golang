package auth

import "context"

func IsAuthenticated(ctx context.Context) error {
	authError := ctx.Value("authError").(error)
	return authError
}