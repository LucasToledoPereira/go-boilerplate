package user_repository

import (
	"errors"

	"github.com/LucasToledoPereira/go-boilerplate/adapters/datastore"
	"github.com/LucasToledoPereira/go-boilerplate/config"
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/LucasToledoPereira/go-boilerplate/internal/enums/codes"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	data datastore.IDatastoreAdapter
}

func NewUserRepository(data datastore.IDatastoreAdapter) (r UserRepo) {
	return UserRepo{
		data: data,
	}
}

func (sr *UserRepo) Create(user *entity.User) error {
	err := sr.data.DB().Omit(clause.Associations).Create(&user).Error
	return err
}

func (sr *UserRepo) Update(user *entity.User) error {
	query := sr.data.DB()
	query = query.Omit(clause.Associations)
	query = query.Save(&user)
	err := query.Error
	return err
}

func (sr *UserRepo) Delete(userID uuid.UUID, studioID uuid.UUID, deleter uuid.UUID) error {
	u := entity.User{ID: userID}
	u.DeletedBy = deleter
	query := sr.data.DB()
	query = query.Where("studio_id = ?", studioID)
	query = query.Delete(&u)
	err := query.Error

	if err == nil && query.RowsAffected == 0 {
		return errors.New(codes.UserNotFound.String())
	}

	return err
}

func (sr *UserRepo) List(studioID uuid.UUID) (users []entity.User, err error) {
	err = sr.data.DB().Where("studio_id = ?", studioID).Find(&users).Error
	return users, err
}

func (sr *UserRepo) Read(id uuid.UUID, studioID uuid.UUID) (user entity.User, err error) {
	query := sr.data.DB()
	query = query.Preload("Studio")
	query = query.First(&user, "id = ? and studio_id = ?", id, studioID)
	err = query.Error

	return user, err
}

func (sr *UserRepo) UpdateImage(user entity.User, key string) error {
	query := sr.data.DB()
	query = query.Omit(clause.Associations)
	query = query.Model(&user)
	query = query.Update("filekey", key)
	return query.Error
}

func (sr *UserRepo) ReadByIdentity(identity string) (user entity.User, err error) {
	query := sr.data.DB()
	query = query.Preload("Studio")
	query = query.Where(config.C.Authorization.IdentityKey+" = ?", identity)
	query = query.Where("exists (select s.id from studios s where s.id = \"users\".\"studio_id\" and s.deleted_at is null)")
	query = query.First(&user)
	err = query.Error
	return
}

func (sr *UserRepo) DeleteByIdentity(identity string, deleter uuid.UUID) (err error) {
	u := &entity.User{}
	u.DeletedBy = deleter
	query := sr.data.DB()
	query = query.Where(config.C.Authorization.IdentityKey+" = ?", identity)
	query = query.Delete(u)
	err = query.Error
	return
}

func (sr *UserRepo) HasUserWithNickname(nickname string) (has bool, err error) {
	query := sr.data.DB()
	query = query.Raw("select true where exists (select u.id from users u where u.nickname = ? and u.deleted_at is null)", nickname)
	query = query.Scan(&has)
	err = query.Error
	return
}

func (sr *UserRepo) HasUserWithEmail(email string) (has bool, err error) {
	query := sr.data.DB()
	query = query.Raw("select true where exists (select u.id from users u where u.email = ? and u.deleted_at is null)", email)
	query = query.Scan(&has)
	err = query.Error
	return
}
