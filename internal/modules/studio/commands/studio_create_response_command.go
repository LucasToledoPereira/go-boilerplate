package studio_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/google/uuid"
)

type CreateStudioCommandResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"fescription"`
	Website     string    `json:"website"`
	Instagram   string    `json:"instagram"`
	CreatedAt   time.Time `json:"created_at"`
}

func (c *CreateStudioCommandResponse) FromEntity(e *entity.Studio) {
	c.ID = e.ID
	c.Name = e.Name
	c.Description = e.Description
	c.Website = e.Website
	c.Instagram = e.Instagram
	c.CreatedAt = e.CreatedAt
}
