package game_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/google/uuid"
)

// command Request to create a user
type CreateGameCommandResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	CreatedAt   time.Time `json:"created_at"`
}

func (c *CreateGameCommandResponse) FromEntity(g *entity.Game) {
	c.ID = g.ID
	c.Title = g.Title
	c.Description = g.Description
	c.Website = g.Website
	c.CreatedAt = g.CreatedAt
}
