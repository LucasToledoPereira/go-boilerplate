package collection_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/google/uuid"
)

// command Request to create a user
type UpdateCollectionAttributeCommandResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Value        string    `json:"value"`
	CollectionID uuid.UUID `json:"collection_id"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    uuid.UUID `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    uuid.UUID `json:"updated_by"`
}

func (c *UpdateCollectionAttributeCommandResponse) FromEntity(tp *entity.CollectionAttributes) {
	c.ID = tp.ID
	c.Name = tp.Name
	c.Type, _ = tp.Type.Value()
	c.Value = tp.Value
	c.CollectionID = tp.CollectionID
	c.CreatedAt = tp.CreatedAt
	c.CreatedBy = tp.CreatedBy
	c.UpdatedAt = tp.UpdatedAt
	c.UpdatedBy = tp.UpdatedBy
}
