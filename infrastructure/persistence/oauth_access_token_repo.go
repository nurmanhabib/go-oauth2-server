package persistence

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"gorm.io/gorm"
)

type OAuthAccessTokenRepo struct {
	db *gorm.DB
}

func NewOAuthAccessTokenRepo(db *gorm.DB) repository.OAuthAccessTokenRepo {
	return &OAuthAccessTokenRepo{db: db}
}

func (o *OAuthAccessTokenRepo) FindByID(ctx context.Context, id string) (*entity.OAuthAccessToken, error) {
	var access entity.OAuthAccessToken

	err := o.db.WithContext(ctx).Take(&access, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &access, nil
}

func (o *OAuthAccessTokenRepo) FindByToken(ctx context.Context, token string) (*entity.OAuthAccessToken, error) {
	var access entity.OAuthAccessToken

	err := o.db.WithContext(ctx).Take(&access, "token = ?", token).Error
	if err != nil {
		return nil, err
	}

	return &access, nil
}

func (o *OAuthAccessTokenRepo) FindByRefreshToken(ctx context.Context, token string) (*entity.OAuthAccessToken, error) {
	var access entity.OAuthAccessToken

	err := o.db.WithContext(ctx).Take(&access, "refresh_token = ?", token).Error
	if err != nil {
		return nil, err
	}

	return &access, nil
}

func (o *OAuthAccessTokenRepo) Save(ctx context.Context, token *entity.OAuthAccessToken) error {
	return o.db.WithContext(ctx).Create(&token).Error
}

func (o *OAuthAccessTokenRepo) Update(ctx context.Context, token *entity.OAuthAccessToken, id string) error {
	q := o.db.WithContext(ctx).Where("id = ?", id).Updates(&token)
	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (o *OAuthAccessTokenRepo) Delete(ctx context.Context, id string) error {
	q := o.db.WithContext(ctx).Where("id = ?", id).Updates(&entity.OAuthAccessToken{})
	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
