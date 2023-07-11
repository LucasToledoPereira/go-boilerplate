package template_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/google/uuid"
)

type ListTemplateAttributeCommandResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Value      string    `json:"value"`
	TemplateID uuid.UUID `json:"template_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (c *ListTemplateAttributeCommandResponse) FromEntity(tp entity.TemplateAttributes) {
	c.ID = tp.ID
	c.Name = tp.Name
	c.Type, _ = tp.Type.Value()
	c.TemplateID = tp.TemplateID
	c.Value = tp.Value
	c.CreatedAt = tp.CreatedAt
	c.UpdatedAt = tp.UpdatedAt
}
