package auth_interfaces

import (
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
)

type IAuthRepository interface {
	Register(user *entity.User) error
	ReadByEmailOrNickname(name string) (user entity.User, err error)
	HasUserWithEmailOrNickname(email string, nickname string) (exists bool, err error)
	HasStudioWithName(name string) (exists bool, err error)
}
