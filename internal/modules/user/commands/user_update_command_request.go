package user_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateUserCommandRequest struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	FullName string `json:"name"`
	Updater  uuid.UUID
}

func (c *UpdateUserCommandRequest) ToEntity(user *entity.User) {
	if c.Email != "" {
		user.Email = c.Email
	}
	if c.Nickname != "" {
		user.Nickname = c.Nickname
	}
	if c.FullName != "" {
		user.FullName = c.FullName
	}
	user.UpdatedBy = c.Updater
}

func (c *UpdateUserCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}
