package store

import (
	"context"
	"fmt"

	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/google/uuid"
	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/infrastructure/dao"
	"gorm.io/gorm"
)

type TokenStore struct {
	repo *dao.Repositories
}

func NewTokenStore(repo *dao.Repositories) oauth2.TokenStore {
	return &TokenStore{repo: repo}
}

func (t *TokenStore) Create(ctx context.Context, info oauth2.TokenInfo) error {
	switch {
	case info.GetCode() != "":
		return t.repo.OAuthAccessGrantRepo.Save(ctx, &entity.OAuthAccessGrant{
			ID:            uuid.New().String(),
			UserID:        info.GetUserID(),
			OAuthClientID: info.GetClientID(),
			Code:          info.GetCode(),
			RedirectURI:   info.GetRedirectURI(),
			Scopes:        info.GetScope(),
			ExpiresIn:     info.GetCodeExpiresIn(),
		})

	case info.GetAccess() != "":
		return t.repo.OAuthAccessTokenRepo.Save(ctx, &entity.OAuthAccessToken{
			ID:               uuid.New().String(),
			UserID:           info.GetUserID(),
			OAuthClientID:    info.GetClientID(),
			Token:            info.GetAccess(),
			Scopes:           info.GetScope(),
			ExpiresIn:        info.GetCodeExpiresIn(),
			RefreshToken:     info.GetRefresh(),
			RefreshExpiresIn: info.GetRefreshExpiresIn(),
		})

	default:
		return nil
	}
}

func (t *TokenStore) RemoveByCode(ctx context.Context, code string) error {
	token, err := t.repo.OAuthAccessGrantRepo.FindByCode(ctx, code)
	if err != nil {
		return err
	}

	return t.repo.OAuthAccessGrantRepo.Delete(ctx, token.ID)
}

func (t *TokenStore) RemoveByAccess(ctx context.Context, access string) error {
	token, err := t.repo.OAuthAccessTokenRepo.FindByToken(ctx, access)
	if err != nil {
		return err
	}

	return t.repo.OAuthAccessTokenRepo.Delete(ctx, token.ID)
}

func (t *TokenStore) RemoveByRefresh(ctx context.Context, refresh string) error {
	token, err := t.repo.OAuthAccessTokenRepo.FindByRefreshToken(ctx, refresh)
	if err != nil {
		return err
	}

	return t.repo.OAuthAccessTokenRepo.Delete(ctx, token.ID)
}

func (t *TokenStore) GetByCode(ctx context.Context, code string) (oauth2.TokenInfo, error) {
	access, err := t.repo.OAuthAccessGrantRepo.FindByCode(ctx, code)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.ErrInvalidAuthorizeCode

		default:
			return nil, fmt.Errorf("failed query by authorize code [%s]: %w", code, err)
		}
	}

	info := models.NewToken()
	info.SetClientID(access.OAuthClientID)
	info.SetUserID(access.UserID)
	info.SetCode(access.Code)
	info.SetCodeCreateAt(access.CreatedAt)
	info.SetCodeExpiresIn(access.ExpiresIn)
	info.SetRedirectURI(access.RedirectURI)
	info.SetScope(access.Scopes)

	return info, nil
}

func (t *TokenStore) GetByAccess(ctx context.Context, access string) (oauth2.TokenInfo, error) {
	token, err := t.repo.OAuthAccessTokenRepo.FindByToken(ctx, access)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.ErrInvalidAccessToken

		default:
			return nil, fmt.Errorf("failed query by access token [%s]: %w", access, err)
		}
	}

	info := models.NewToken()
	info.SetClientID(token.OAuthClientID)
	info.SetUserID(token.UserID)
	info.SetAccess(token.Token)
	info.SetAccessCreateAt(token.CreatedAt)
	info.SetAccessExpiresIn(token.ExpiresIn)
	info.SetRefresh(token.RefreshToken)
	info.SetRefreshExpiresIn(token.RefreshExpiresIn)
	info.SetRefreshCreateAt(token.CreatedAt)
	info.SetScope(token.Scopes)

	return info, nil
}

func (t *TokenStore) GetByRefresh(ctx context.Context, refresh string) (oauth2.TokenInfo, error) {
	access, err := t.repo.OAuthAccessTokenRepo.FindByRefreshToken(ctx, refresh)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.ErrInvalidRefreshToken

		default:
			return nil, fmt.Errorf("failed query by refresh token [%s]: %w", refresh, err)
		}
	}

	info := models.NewToken()
	info.SetClientID(access.OAuthClientID)
	info.SetUserID(access.UserID)
	info.SetAccess(access.Token)
	info.SetAccessCreateAt(access.CreatedAt)
	info.SetAccessExpiresIn(access.ExpiresIn)
	info.SetRefresh(access.RefreshToken)
	info.SetRefreshExpiresIn(access.RefreshExpiresIn)
	info.SetRefreshCreateAt(access.CreatedAt)
	info.SetScope(access.Scopes)

	return info, nil
}
