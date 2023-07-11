package template_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/google/uuid"
)

// command Request to create a user
type CreateTemplateAttributeCommandResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *CreateTemplateAttributeCommandResponse) FromEntity(tp *entity.TemplateAttributes) {
	c.ID = tp.ID
	c.Name = tp.Name
	c.Type, _ = tp.Type.Value()
	c.Value = tp.Value
	c.CreatedAt = tp.CreatedAt
}
