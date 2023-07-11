package auth_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/internal/utils"
	"github.com/gin-gonic/gin"
)

// command Request to create a user
type LoginCommandRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func (c *LoginCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}

func (c *LoginCommandRequest) Validate() (messages []string) {
	messages = make([]string, 1)

	if c.Nickname == "" {
		messages[0] = "email.or.nickname.required"
		return messages
	}

	if c.Password == "" {
		messages[0] = "password.required"
		return messages
	}

	return nil
}

func (usr *LoginCommandRequest) CheckPassword(password string) bool {
	return utils.EncryptPassword(usr.Password) == password
}
