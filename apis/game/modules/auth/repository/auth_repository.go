package auth_repository

import (
	"errors"

	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AuthRepo struct {
	data gbp.IDatastoreAdapter
}

func NewAuthRepository(data gbp.IDatastoreAdapter) (r AuthRepo) {
	return AuthRepo{
		data: data,
	}
}

func (sr *AuthRepo) Register(user *entity.User) error {
	err := sr.data.DB().Clauses(clause.OnConflict{DoNothing: true}).Create(&user).Error
	return err
}

func (sr *AuthRepo) ReadByEmailOrNickname(name string) (user entity.User, err error) {
	err = sr.data.DB().Preload("Studio").Where("email = ?", name).Or("nickname = ?", name).First(&user).Error
	return user, err
}

func (sr *AuthRepo) HasUserWithEmailOrNickname(email string, nickname string) (bool, error) {
	err := sr.data.DB().
		Where("email = ?", email).
		Or("nickname = ?", nickname).
		First(&entity.User{}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, err
}

func (sr *AuthRepo) HasStudioWithName(name string) (bool, error) {
	err := sr.data.DB().
		Where("name = ?", name).
		First(&entity.Studio{}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return true, err
}
