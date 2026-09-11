package auth

import (
	"context"
	"fmt"

	"github.com/AlekSi/pointer"
	v1 "github.com/RenterRus/sausage-auth/docs/proto/v1"
	"github.com/RenterRus/sausage-bff/internal/entity"
)

type auth struct {
	client v1.AuthServiceClient
}

func NewAuthController(client v1.AuthServiceClient) AuthService {
	return &auth{
		client: client,
	}
}

// Confirm implements AuthService.
func (a *auth) Confirm(ctx context.Context, req *entity.AcceptRequest) (string, error) {
	if req == nil {
		return "", fmt.Errorf("Confirm: %w", entity.ErrNoRequiredParams)
	}

	resp, err := a.client.Confirm(ctx, &v1.AcceptRequest{
		Login:   req.Login,
		OtpCode: req.OtpCode,
	})
	if err != nil {
		return "", fmt.Errorf("Confirm.Confirm: %w", err)
	}

	return resp.GetStatus(), nil
}

// LoginOTP implements AuthService.
func (a *auth) LoginOTP(ctx context.Context, req *entity.LoginOTPRequest) (*entity.Tokens, error) {
	if req == nil {
		return nil, fmt.Errorf("LoginOTP: %w", entity.ErrNoRequiredParams)
	}

	resp, err := a.client.LoginOTP(ctx, &v1.LoginOTPRequest{
		Login:     req.Login,
		UserAgent: req.UserAgent,
		Code:      req.Code,
	})
	if err != nil {
		return nil, fmt.Errorf("LoginOTP.LoginOTP: %w", err)
	}

	return &entity.Tokens{
		Access:  resp.GetTokens().GetAccess(),
		Refresh: resp.GetTokens().GetRefresh(),
	}, nil
}

// RefreshToken implements AuthService.
func (a *auth) RefreshToken(ctx context.Context, req *entity.RefreshRequest) (*entity.Tokens, error) {
	if req == nil {
		return nil, fmt.Errorf("RefreshToken: %w", entity.ErrNoRequiredParams)
	}

	resp, err := a.client.RefreshToken(ctx, &v1.RefreshRequest{
		Login:        req.Login,
		UserAgent:    req.UserAgent,
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		return nil, fmt.Errorf("RefreshToken.RefreshToken: %w", err)
	}

	return &entity.Tokens{
		Access:  resp.GetTokens().GetAccess(),
		Refresh: resp.GetTokens().GetRefresh(),
	}, nil
}

// Register implements AuthService.
func (a *auth) Register(ctx context.Context, login *string) (string, error) {
	if pointer.Get(login) == "" {
		return "", fmt.Errorf("Register: %w", entity.ErrNoRequiredParams)
	}

	url, err := a.client.Register(ctx, &v1.RegisterRequest{
		Login: *login,
	})
	if err != nil {
		return "", fmt.Errorf("Register.Register: %w", err)
	}

	return url.GetUrl(), nil
}

// RevokeSession implements AuthService.
func (a *auth) RevokeSession(ctx context.Context, req *entity.RevokeSessionRequest) (string, error) {
	if req == nil {
		return "", fmt.Errorf("RevokeSession: %w", entity.ErrNoRequiredParams)
	}

	var resp *v1.RevokeSessionResponse
	var err error

	if req.Current != nil {
		resp, err = a.client.RevokeSession(ctx, &v1.RevokeSessionRequest{
			Mode: &v1.RevokeSessionRequest_Current_{
				Current: &v1.RevokeSessionRequest_Current{
					Hash: req.Current.Hash,
				},
			},
		})
		if err != nil {
			return "", fmt.Errorf("RevokeSession.Current: %w", err)
		}
	}

	if req.All != nil {
		resp, err = a.client.RevokeSession(ctx, &v1.RevokeSessionRequest{
			Mode: &v1.RevokeSessionRequest_All_{
				All: &v1.RevokeSessionRequest_All{
					Login: req.All.Login,
				},
			},
		})
		if err != nil {
			return "", fmt.Errorf("RevokeSession.All: %w", err)
		}
	}

	return resp.GetStatus(), nil
}

// ValidateToken implements AuthService.
func (a *auth) ValidateToken(ctx context.Context, access *string) (string, error) {
	if pointer.Get(access) == "" {
		return "", fmt.Errorf("ValidateToken: %w", entity.ErrNoRequiredParams)
	}

	resp, err := a.client.ValidateToken(ctx, &v1.ValidateTokenRequest{
		Access: *access,
	})
	if err != nil {
		return "", fmt.Errorf("ValidateToken.ValidateToken: %w", err)
	}

	return resp.GetUuid(), nil
}
