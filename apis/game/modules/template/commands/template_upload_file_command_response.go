package template_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/google/uuid"
)

// command Request to create a user
type UploadFileTemplateCommandResponse struct {
	ID         uuid.UUID `json:"id"`
	TemplateID uuid.UUID `json:"template_id"`
	URL        string    `json:"url"`
	Type       string    `json:"type"`
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  uuid.UUID `json:"created_by"`
	UpdatedAt  time.Time `json:"updated_at"`
	UpdatedBy  uuid.UUID `json:"updated_by"`
}

func (c *UploadFileTemplateCommandResponse) FromEntity(tp entity.TemplateFiles, url string) {
	c.ID = tp.ID
	c.TemplateID = tp.TemplateID
	c.URL = url
	c.Type, _ = tp.Type.Value()
	c.Key = tp.Key
	c.CreatedAt = tp.CreatedAt
	c.CreatedBy = tp.CreatedBy
	c.UpdatedAt = tp.UpdatedAt
	c.UpdatedBy = tp.UpdatedBy
}
