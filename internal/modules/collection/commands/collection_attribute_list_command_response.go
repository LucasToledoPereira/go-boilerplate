package collection_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/google/uuid"
)

type ListCollectionAttributesCommandResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Value        string    `json:"value"`
	CollectionID uuid.UUID `json:"collection_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (c *ListCollectionAttributesCommandResponse) FromEntity(tp entity.CollectionAttributes) {
	c.ID = tp.ID
	c.Name = tp.Name
	c.Type, _ = tp.Type.Value()
	c.CollectionID = tp.CollectionID
	c.Value = tp.Value
	c.CreatedAt = tp.CreatedAt
	c.UpdatedAt = tp.UpdatedAt
}
