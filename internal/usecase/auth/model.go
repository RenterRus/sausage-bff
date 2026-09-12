package auth

import (
	"context"

	"github.com/RenterRus/sausage-bff/internal/entity"
)

type AuthCases interface {
	Register(ctx context.Context, login *string) (string, error)
	Confirm(ctx context.Context, req *entity.AcceptRequest) (string, error)
	LoginOTP(ctx context.Context, req *entity.LoginOTPRequest) (*entity.Tokens, error)
	UrlOTP(ctx context.Context, access *string) (string, error)
	RevokeSession(ctx context.Context, req *entity.RevokeSessionRequest) (string, error)
	ValidateToken(ctx context.Context, access *string) (string, error)
	RefreshToken(ctx context.Context, req *entity.RefreshRequest) (*entity.Tokens, error)
}
