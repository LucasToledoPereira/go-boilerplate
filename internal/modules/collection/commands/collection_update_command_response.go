package collection_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/google/uuid"
)

// command Request to create a user
type UpdateCollectionCommandResponse struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Symbol           string    `json:"symbol"`
	Description      string    `json:"description"`
	ShortDescription string    `json:"short_description"`
	CreatedAt        time.Time `json:"created_at"`
	GameID           uuid.UUID `json:"game_id"`
}

func (c *UpdateCollectionCommandResponse) FromEntity(tp *entity.Collection) {
	c.ID = tp.ID
	c.Name = tp.Name
	c.Symbol = tp.Symbol
	c.Description = tp.Description
	c.ShortDescription = tp.ShortDescription
	c.CreatedAt = tp.CreatedAt
	c.GameID = tp.GameID
}
