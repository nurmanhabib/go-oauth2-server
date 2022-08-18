package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/google/uuid"
	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"github.com/nurmanhabib/go-oauth2-server/pkg/oauth/config"
)

// TokenStorage is a container for token storage needs.
type TokenStorage struct {
	accessGrant repository.OauthAccessGrantRepository
	accessToken repository.OauthAccessTokenRepository
	client      repository.OauthClientRepository
	user        repository.UserRepository
}

// WithTokenStorage is a function to set to OAuth config.
func WithTokenStorage(
	accessGrant repository.OauthAccessGrantRepository,
	accessToken repository.OauthAccessTokenRepository,
	client repository.OauthClientRepository,
	user repository.UserRepository,
) config.Option {
	return func(c *config.Config) {
		c.TokenStorage = NewTokenStorage(accessGrant, accessToken, client, user)
	}
}

// NewTokenStorage is constructor.
func NewTokenStorage(
	accessGrant repository.OauthAccessGrantRepository,
	accessToken repository.OauthAccessTokenRepository,
	client repository.OauthClientRepository,
	user repository.UserRepository,
) *TokenStorage {
	return &TokenStorage{
		accessGrant: accessGrant,
		accessToken: accessToken,
		client:      client,
		user:        user,
	}
}

// Create is a function to store tokens.
func (t *TokenStorage) Create(ctx context.Context, info oauth2.TokenInfo) error {
	if info.GetCode() != "" {
		return t.createAccessGrant(ctx, info)
	}

	return t.createAccessToken(ctx, info)
}

func (t *TokenStorage) createAccessGrant(ctx context.Context, info oauth2.TokenInfo) error {
	user, err := t.user.Find(ctx, &entity.User{ID: info.GetUserID()})
	if err != nil {
		return err
	}

	client, err := t.client.Find(ctx, &entity.OauthClient{ID: info.GetClientID()})
	if err != nil {
		return err
	}

	err = t.accessGrant.Save(ctx, &entity.OauthAccessGrant{
		ID:            uuid.New().String(),
		UserID:        user.ID,
		OauthClientID: client.ID,
		Code:          info.GetCode(),
		RedirectURI:   info.GetRedirectURI(),
		ExpiresIn:     int(info.GetCodeExpiresIn().Seconds()),
		Scopes:        info.GetScope(),
		RevokedAt:     sql.NullTime{},
		CreatedAt:     info.GetCodeCreateAt(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (t *TokenStorage) createAccessToken(ctx context.Context, info oauth2.TokenInfo) error {
	var err error
	var user *entity.User
	var userID sql.NullString

	user = &entity.User{}

	// If the User ID is empty, then the access token belongs to the client (application).
	if info.GetUserID() != "" {
		user, err = t.user.Find(ctx, &entity.User{ID: info.GetUserID()})
		if err != nil {
			return err
		}

		_ = userID.Scan(user.ID)
	}

	client, err := t.client.Find(ctx, &entity.OauthClient{ID: info.GetClientID()})
	if err != nil {
		return err
	}

	err = t.accessToken.Save(ctx, &entity.OauthAccessToken{
		ID:            uuid.New().String(),
		UserID:        userID,
		OauthClientID: client.ID,
		Token:         info.GetAccess(),
		RefreshToken:  info.GetRefresh(),
		ExpiresIn:     int(info.GetAccessExpiresIn().Seconds()),
		Scopes:        info.GetScope(),
		RevokedAt:     sql.NullTime{},
		CreatedAt:     info.GetAccessCreateAt(),
	})
	if err != nil {
		return err
	}

	return nil
}

// RemoveByCode is a function for revoke authorization code.
func (t *TokenStorage) RemoveByCode(ctx context.Context, code string) error {
	err := t.accessGrant.Revoke(ctx, &entity.OauthAccessGrant{Code: code}, time.Now())
	if err != nil {
		return err
	}

	return nil
}

// RemoveByAccess is a function for revoke access token.
func (t *TokenStorage) RemoveByAccess(ctx context.Context, access string) error {
	err := t.accessToken.Revoke(ctx, &entity.OauthAccessToken{Token: access}, time.Now())
	if err != nil {
		return err
	}

	return nil
}

// RemoveByRefresh is a function for revoke refresh token.
func (t *TokenStorage) RemoveByRefresh(ctx context.Context, refresh string) error {
	err := t.accessToken.Revoke(ctx, &entity.OauthAccessToken{RefreshToken: refresh}, time.Now())
	if err != nil {
		return err
	}

	return nil
}

// GetByCode is a function to retrieve the authorization code.
func (t *TokenStorage) GetByCode(ctx context.Context, code string) (oauth2.TokenInfo, error) {
	accessGrant, err := t.accessGrant.Find(ctx, &entity.OauthAccessGrant{Code: code})
	if err != nil {
		return nil, err
	}

	if accessGrant.RevokedAt.Valid {
		return nil, errors.ErrInvalidGrant
	}

	client, err := t.client.Find(ctx, &entity.OauthClient{ID: accessGrant.OauthClientID})
	if err != nil {
		return nil, err
	}

	user, err := t.user.Find(ctx, &entity.User{ID: accessGrant.UserID})
	if err != nil {
		return nil, err
	}

	tokenInfo := &models.Token{
		ClientID:      client.ID,
		UserID:        user.ID,
		RedirectURI:   accessGrant.RedirectURI,
		Scope:         accessGrant.Scopes,
		Code:          accessGrant.Code,
		CodeCreateAt:  accessGrant.CreatedAt,
		CodeExpiresIn: time.Duration(accessGrant.ExpiresIn * int(time.Second)),
	}

	return tokenInfo, nil
}

// GetByAccess is a function to retrieve the access token.
func (t *TokenStorage) GetByAccess(ctx context.Context, access string) (oauth2.TokenInfo, error) {
	return t.getByAccessToken(ctx, &entity.OauthAccessToken{Token: access})
}

// GetByRefresh is a function to retrieve the refresh token.
func (t *TokenStorage) GetByRefresh(ctx context.Context, refresh string) (oauth2.TokenInfo, error) {
	return t.getByAccessToken(ctx, &entity.OauthAccessToken{RefreshToken: refresh})
}

func (t *TokenStorage) getByAccessToken(ctx context.Context, find *entity.OauthAccessToken) (oauth2.TokenInfo, error) {
	accessToken, err := t.accessToken.Find(ctx, find)
	if err != nil {
		return nil, err
	}

	client, err := t.client.Find(ctx, &entity.OauthClient{ID: accessToken.OauthClientID})
	if err != nil {
		return nil, err
	}

	user, err := t.user.Find(ctx, &entity.User{ID: accessToken.UserID.String})
	if err != nil {
		return nil, err
	}

	tokenInfo := &models.Token{
		ClientID:         client.ID,
		UserID:           user.ID,
		Scope:            accessToken.Scopes,
		Access:           accessToken.Token,
		AccessCreateAt:   accessToken.CreatedAt,
		AccessExpiresIn:  time.Duration(accessToken.ExpiresIn * int(time.Second)),
		Refresh:          accessToken.RefreshToken,
		RefreshCreateAt:  accessToken.CreatedAt,
		RefreshExpiresIn: time.Duration(accessToken.RefreshExpiresIn * int(time.Second)),
	}

	return tokenInfo, nil
}
