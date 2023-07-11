package studio_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateStudioCommandRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Instagram   string `json:"instagram"`
}

func (c *CreateStudioCommandRequest) ToEntity() (studio *entity.Studio) {
	return &entity.Studio{
		ID:          uuid.New(),
		Name:        c.Name,
		Description: c.Description,
		Website:     c.Website,
		Instagram:   c.Instagram,
	}
}

func (c *CreateStudioCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}

func (c *CreateStudioCommandRequest) Validate() (messages []string) {
	if c.Name == "" {
		messages = make([]string, 1)
		messages[0] = "name.required"
		return messages
	}
	return nil
}
