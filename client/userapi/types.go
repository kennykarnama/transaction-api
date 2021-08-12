package userapi

import "context"

type ValidateTokenRequest struct {
}

type Client interface {
	ValidateToken(ctx context.Context, token string) error
}
