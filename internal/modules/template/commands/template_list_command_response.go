package template_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/google/uuid"
)

type ListTemplateCommandResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *ListTemplateCommandResponse) FromEntity(tp entity.Template) {
	c.ID = tp.ID
	c.Name = tp.Name
	c.Symbol = tp.Symbol
	c.CreatedAt = tp.CreatedAt
	c.UpdatedAt = tp.UpdatedAt
}

func ListResponseFromEntities(templates []entity.Template) (res []ListTemplateCommandResponse) {
	for _, template := range templates {
		var lsc ListTemplateCommandResponse
		lsc.FromEntity(template)
		res = append(res, lsc)
	}
	return res
}
