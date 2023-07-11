package template_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/google/uuid"
)

// command Request to create a user
type UpdateTemplateCommandResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Symbol       string    `json:"symbol"`
	Description  string    `json:"description"`
	AnimationURL string    `json:"animation_url"`
	ExternalURL  string    `json:"external_url"`
	Category     string    `json:"category"`
	Supply       string    `json:"supply"`
	CreatedAt    time.Time `json:"created_at"`
	GameID       uuid.UUID `json:"game_id"`
}

func (c *UpdateTemplateCommandResponse) FromEntity(tp *entity.Template) {
	c.ID = tp.ID
	c.Name = tp.Name
	c.Symbol = tp.Symbol
	c.Description = tp.Description
	c.AnimationURL = tp.AnimationURL
	c.ExternalURL = tp.ExternalURL
	c.Category = tp.Category
	c.Supply = tp.Supply
	c.CreatedAt = tp.CreatedAt
	c.GameID = tp.GameID
}
