package game_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateGameCommandRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Updater     uuid.UUID
	StudioID    uuid.UUID
	GameID      uuid.UUID
}

func (c *UpdateGameCommandRequest) ToEntity(game *entity.Game) {
	if c.Title != "" {
		game.Title = c.Title
	}
	game.Description = c.Description
	game.Website = c.Website
	game.UpdatedBy = c.Updater
}

func (c *UpdateGameCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}
