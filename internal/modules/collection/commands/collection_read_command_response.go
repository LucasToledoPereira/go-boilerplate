package collection_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/google/uuid"
)

// command Request to create a user
type ReadCollectionCommandResponse struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Symbol           string    `json:"symbol"`
	Description      string    `json:"description"`
	ShortDescription string    `json:"short_description"`
	Address          string    `json:"address"`
	Cover            string    `json:"cover"`
	Avatar           string    `json:"avatar"`
	CreatedAt        time.Time `json:"created_at"`
	Game             struct {
		ID    uuid.UUID `json:"id"`
		Title string    `json:"title"`
	} `json:"game"`
}

func (c *ReadCollectionCommandResponse) FromEntity(tp *entity.Collection, cover, avatar string) {
	c.ID = tp.ID
	c.Name = tp.Name
	c.Symbol = tp.Symbol
	c.Description = tp.Description
	c.ShortDescription = tp.ShortDescription
	c.Address = tp.Address
	c.CreatedAt = tp.CreatedAt
	c.Avatar = avatar
	c.Cover = cover
	c.Game.ID = tp.Game.ID
	c.Game.Title = tp.Game.Title
}
