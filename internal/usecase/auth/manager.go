package auth

import (
	"context"
	"fmt"

	controllerAuth "github.com/RenterRus/sausage-bff/internal/controller/grpc/auth"
	"github.com/RenterRus/sausage-bff/internal/entity"
)

type authCase struct {
	api controllerAuth.AuthService
}

func NewAuthCase(ctrl controllerAuth.AuthService) AuthCases {
	return &authCase{
		api: ctrl,
	}
}

// Confirm implements AuthCases.
func (a *authCase) Confirm(ctx context.Context, req *entity.AcceptRequest) (string, error) {
	if req == nil {
		return "", fmt.Errorf("Confirm: %w", entity.ErrNoRequiredParams)
	}

	uuid, err := a.api.Confirm(ctx, req)
	if err != nil {
		return "", fmt.Errorf("Confirm.Confirm: %w", err)
	}

	return uuid, nil
}

// LoginOTP implements AuthCases.
func (a *authCase) LoginOTP(ctx context.Context, req *entity.LoginOTPRequest) (*entity.Tokens, error) {
	if req == nil {
		return nil, fmt.Errorf("LoginOTP: %w", entity.ErrNoRequiredParams)
	}

	resp, err := a.api.LoginOTP(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("LoginOTP.LoginOTP: %w", err)
	}

	return resp, nil
}

// RefreshToken implements AuthCases.
func (a *authCase) RefreshToken(ctx context.Context, req *entity.RefreshRequest) (*entity.Tokens, error) {
	if req == nil {
		return nil, fmt.Errorf("RefreshToken: %w", entity.ErrNoRequiredParams)
	}

	resp, err := a.api.RefreshToken(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("RefreshToken.RefreshToken: %w", err)
	}

	return resp, nil
}

// Register implements AuthCases.
func (a *authCase) Register(ctx context.Context, login *string) (string, error) {
	if login == nil {
		return "", fmt.Errorf("Register: %w", entity.ErrNoRequiredParams)
	}

	url, err := a.api.Register(ctx, login)
	if err != nil {
		return "", fmt.Errorf("RefreshToken.RefreshToken: %w", err)
	}

	return url, nil
}

// RevokeSession implements AuthCases.
func (a *authCase) RevokeSession(ctx context.Context, req *entity.RevokeSessionRequest) (string, error) {
	if req == nil {
		return "", fmt.Errorf("RevokeSession: %w", entity.ErrNoRequiredParams)
	}

	resp, err := a.api.RevokeSession(ctx, req)
	if err != nil {
		return "", fmt.Errorf("RevokeSession.RevokeSession: %w", err)
	}

	return resp, nil
}

// UrlOTP implements AuthCases.
func (a *authCase) UrlOTP(ctx context.Context, access *string) (string, error) {
	if access == nil {
		return "", fmt.Errorf("UrlOTP: %w", entity.ErrNoRequiredParams)
	}

	url, err := a.api.UrlOTP(ctx, access)
	if err != nil {
		return "", fmt.Errorf("UrlOTP.UrlOTP: %w", err)
	}

	return url, nil
}

// ValidateToken implements AuthCases.
func (a *authCase) ValidateToken(ctx context.Context, access *string) (string, error) {
	if access == nil {
		return "", fmt.Errorf("ValidateToken: %w", entity.ErrNoRequiredParams)
	}

	uuid, err := a.api.ValidateToken(ctx, access)
	if err != nil {
		return "", fmt.Errorf("ValidateToken.ValidateToken: %w", err)
	}

	return uuid, nil
}
