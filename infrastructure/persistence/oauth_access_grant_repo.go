package persistence

import (
	"context"

	"github.com/nurmanhabib/go-oauth2-server/domain/entity"
	"github.com/nurmanhabib/go-oauth2-server/domain/repository"
	"gorm.io/gorm"
)

type OAuthAccessGrantRepo struct {
	db *gorm.DB
}

func NewOAuthAccessGrantRepo(db *gorm.DB) repository.OAuthAccessGrantRepo {
	return &OAuthAccessGrantRepo{db: db}
}

func (o *OAuthAccessGrantRepo) FindByID(ctx context.Context, id string) (*entity.OAuthAccessGrant, error) {
	var grant entity.OAuthAccessGrant

	err := o.db.WithContext(ctx).Take(&grant, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &grant, nil
}

func (o *OAuthAccessGrantRepo) FindByCode(ctx context.Context, code string) (*entity.OAuthAccessGrant, error) {
	var grant entity.OAuthAccessGrant

	err := o.db.WithContext(ctx).Take(&grant, "code = ?", code).Error
	if err != nil {
		return nil, err
	}

	return &grant, nil
}

func (o *OAuthAccessGrantRepo) Save(ctx context.Context, grant *entity.OAuthAccessGrant) error {
	return o.db.WithContext(ctx).Create(&grant).Error
}

func (o *OAuthAccessGrantRepo) Update(ctx context.Context, grant *entity.OAuthAccessGrant, id string) error {
	q := o.db.WithContext(ctx).Where("id = ?", id).Updates(&grant)
	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (o *OAuthAccessGrantRepo) Delete(ctx context.Context, id string) error {
	q := o.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.OAuthAccessGrant{})
	if q.Error != nil {
		return q.Error
	}

	if q.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
