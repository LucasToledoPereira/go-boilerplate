package studio_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/gin-gonic/gin"
)

type UpdateStudioCommandRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Instagram   string `json:"instagram"`
}

func (c *UpdateStudioCommandRequest) ToEntity(studio *entity.Studio) {
	if c.Name != "" {
		studio.Name = c.Name
	}
	studio.Description = c.Description
	studio.Website = c.Website
	studio.Instagram = c.Instagram
}

func (c *UpdateStudioCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}

func (c *UpdateStudioCommandRequest) Validate() (messages []string) {
	if c.Name == "" {
		messages = make([]string, 1)
		messages[0] = "name.required"
		return messages
	}
	return nil
}
