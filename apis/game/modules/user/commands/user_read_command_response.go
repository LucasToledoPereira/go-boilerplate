package user_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/google/uuid"
)

// command Request to create a user
type ReadUserCommandResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Type      string    `json:"type"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	Studio    struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
	} `json:"studio"`
}

func (c *ReadUserCommandResponse) FromEntity(u *entity.User, fileTempURL string) {
	c.ID = u.ID
	c.Name = u.FullName
	c.Email = u.Email
	c.Nickname = u.Nickname
	c.Type, _ = u.Type.Value()
	c.CreatedAt = u.CreatedAt
	c.Studio.Name = u.Studio.Name
	c.Studio.Description = u.Studio.Description
	c.Studio.ID = u.Studio.ID
	c.Image = fileTempURL
}
