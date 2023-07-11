package user_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/google/uuid"
)

type ListUserCommandResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Type      string    `json:"type"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func (c *ListUserCommandResponse) FromEntity(e entity.User, url string) {
	role, _ := e.Type.Value()
	c.ID = e.ID
	c.Name = e.FullName
	c.Email = e.Email
	c.Nickname = e.Nickname
	c.Type = role
	c.CreatedAt = e.CreatedAt
	c.UpdatedAt = e.UpdatedAt
	c.Image = url
}
