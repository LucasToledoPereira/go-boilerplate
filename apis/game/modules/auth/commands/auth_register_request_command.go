package auth_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/LucasToledoPereira/go-boilerplate/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RegisterCommandRequest struct {
	Email      string `json:"email"`
	Nickname   string `json:"nickname"`
	FullName   string `json:"fullname"`
	Password   string `json:"password"`
	StudioName string `json:"studio_name"`
}

func (c *RegisterCommandRequest) ToEntity() (studio *entity.User) {
	return &entity.User{
		ID:       uuid.New(),
		Email:    c.Email,
		Nickname: c.Nickname,
		FullName: c.FullName,
		Password: utils.EncryptPassword(c.Password),
		Type:     datatypes.OWNER,
		Studio: entity.Studio{
			ID:   uuid.New(),
			Name: c.StudioName,
		},
	}
}

func (c *RegisterCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}

func (c *RegisterCommandRequest) Validate() (messages []string) {
	messages = make([]string, 1)

	if c.Email == "" {
		messages[0] = "email.required"
		return messages
	}

	if c.Nickname == "" {
		messages[0] = "nickname.required"
		return messages
	}

	if c.Password == "" {
		messages[0] = "password.required"
		return messages
	}

	if c.FullName == "" {
		messages[0] = "fullname.required"
		return messages
	}

	if c.StudioName == "" {
		messages[0] = "studioname.required"
		return messages
	}

	return nil
}
