package studio_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/google/uuid"
)

type ReadStudioCommandResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	Instagram   string    `json:"instagram"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
}

func (c *ReadStudioCommandResponse) FromEntity(e *entity.Studio, url string) {
	c.ID = e.ID
	c.Name = e.Name
	c.Description = e.Description
	c.Website = e.Website
	c.Instagram = e.Instagram
	c.CreatedAt = e.CreatedAt
	c.UpdatedAt = e.UpdatedAt
	c.Image = url
}
