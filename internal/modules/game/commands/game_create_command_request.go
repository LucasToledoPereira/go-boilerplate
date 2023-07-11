package game_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateGameCommandRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Creator     uuid.UUID
	StudioID    uuid.UUID
}

func (c *CreateGameCommandRequest) ToEntity() entity.Game {
	g := entity.Game{
		ID:          uuid.New(),
		Title:       c.Title,
		Description: c.Description,
		Website:     c.Website,
		StudioID:    c.StudioID,
	}
	g.CreatedBy = c.Creator
	return g
}

func (c *CreateGameCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}

func (c *CreateGameCommandRequest) Validate() (messages []string) {
	if c.Title == "" {
		messages = make([]string, 1)
		messages[0] = "title.required"
		return messages
	}
	return nil
}
