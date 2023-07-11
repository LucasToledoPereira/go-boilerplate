package collection_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/google/uuid"
)

type ListCollectionCommandResponse struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Symbol           string    `json:"symbol"`
	Description      string    `json:"description"`
	ShortDescription string    `json:"short_description"`
	Address          string    `json:"address"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (c *ListCollectionCommandResponse) FromEntity(tp entity.Collection) {
	c.ID = tp.ID
	c.Name = tp.Name
	c.Symbol = tp.Symbol
	c.Description = tp.Description
	c.ShortDescription = tp.ShortDescription
	c.Address = tp.Address
	c.CreatedAt = tp.CreatedAt
	c.UpdatedAt = tp.UpdatedAt
}
