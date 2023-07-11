package template_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/google/uuid"
)

// command Request to create a user
type ReadTemplateCommandResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Symbol       string    `json:"symbol"`
	Description  string    `json:"description"`
	AnimationURL string    `json:"animation_url"`
	ExternalURL  string    `json:"external_url"`
	Category     string    `json:"category"`
	Supply       string    `json:"supply"`
	CreatedAt    time.Time `json:"created_at"`
	Game         struct {
		ID    uuid.UUID `json:"id"`
		Title string    `json:"title"`
	} `json:"game"`
}

func (c *ReadTemplateCommandResponse) FromEntity(tp *entity.Template) {
	c.ID = tp.ID
	c.Name = tp.Name
	c.Symbol = tp.Symbol
	c.Description = tp.Description
	c.AnimationURL = tp.AnimationURL
	c.ExternalURL = tp.ExternalURL
	c.Category = tp.Category
	c.Supply = tp.Supply
	c.CreatedAt = tp.CreatedAt
	c.Game.ID = tp.Game.ID
	c.Game.Title = tp.Game.Title
}
