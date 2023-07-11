package game_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/google/uuid"
)

// command Request to create a user
type UpdateGameCommandResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   uuid.UUID `json:"created_by"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	Studio      struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
	} `json:"studio"`
}

func (c *UpdateGameCommandResponse) FromEntity(g *entity.Game) {
	c.ID = g.ID
	c.Title = g.Title
	c.Description = g.Description
	c.Website = g.Website
	c.CreatedAt = g.CreatedAt
	c.UpdatedAt = g.UpdatedAt
	c.CreatedBy = g.CreatedBy
	c.UpdatedBy = g.UpdatedBy
	c.Studio.Name = g.Studio.Name
	c.Studio.Description = g.Studio.Description
	c.Studio.ID = g.Studio.ID
}
