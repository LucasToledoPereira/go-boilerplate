package user_interfaces

import (
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/google/uuid"
)

type IUserRepository interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(userID uuid.UUID, studioID uuid.UUID, deleter uuid.UUID) error
	List(studioID uuid.UUID) (users []entity.User, err error)
	Read(id uuid.UUID, studioID uuid.UUID) (user entity.User, err error)
	ReadByIdentity(identity string) (user entity.User, err error)
	DeleteByIdentity(identity string, deleter uuid.UUID) (err error)
	HasUserWithNickname(nickname string) (bool, error)
	HasUserWithEmail(email string) (bool, error)
	UpdateImage(user entity.User, filekey string) error
}
