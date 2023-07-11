package user_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/internal/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/LucasToledoPereira/go-boilerplate/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserCommandRequest struct {
	Email    string             `json:"email"`
	Nickname string             `json:"nickname"`
	FullName string             `json:"name"`
	Password string             `json:"password"`
	Type     datatypes.UserRole `json:"type"`
	Creator  uuid.UUID
}

func (c *CreateUserCommandRequest) ToEntity(studioID uuid.UUID) entity.User {
	u := entity.User{
		ID:       uuid.New(),
		Nickname: c.Nickname,
		FullName: c.FullName,
		Email:    c.Email,
		Password: utils.EncryptPassword(c.Password),
		Type:     c.Type,
		StudioID: studioID,
	}
	u.CreatedBy = c.Creator
	return u
}

func (c *CreateUserCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}

func (c *CreateUserCommandRequest) Validate() (messages []string) {
	if c.FullName == "" {
		messages = make([]string, 1)
		messages[0] = "name.required"
		return messages
	}

	if c.Nickname == "" {
		messages = make([]string, 1)
		messages[0] = "nickname.required"
		return messages
	}

	if c.Email == "" {
		messages = make([]string, 1)
		messages[0] = "email.required"
		return messages
	}

	if c.Password == "" {
		messages = make([]string, 1)
		messages[0] = "password.required"
		return messages
	}

	if c.Type.IsNotValid() {
		messages = make([]string, 1)
		messages[0] = "type.invalid"
		return messages
	}
	return nil
}
