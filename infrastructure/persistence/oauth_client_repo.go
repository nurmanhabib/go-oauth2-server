package persistence

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"gorm.io/gorm"
)

type OAuthClientRepo struct {
	db *gorm.DB
}

func NewOAuthClientRepo(db *gorm.DB) repository.OAuthClientRepo {
	return &OAuthClientRepo{db: db}
}

func (o *OAuthClientRepo) FindByID(ctx context.Context, id string) (*entity.OAuthClient, error) {
	var model entity.OAuthClient

	err := o.db.WithContext(ctx).Take(&model, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (o *OAuthClientRepo) Save(ctx context.Context, client *entity.OAuthClient) error {
	return o.db.WithContext(ctx).Create(&client).Error
}

func (o *OAuthClientRepo) Update(ctx context.Context, client *entity.OAuthClient, id string) error {
	q := o.db.WithContext(ctx).Where("id = ?", id).Updates(&client)
	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (o *OAuthClientRepo) Delete(ctx context.Context, id string) error {
	q := o.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.OAuthClient{})
	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
