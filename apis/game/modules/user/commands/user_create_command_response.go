package user_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/google/uuid"
)

// command Request to create a user
type CreateUserCommandResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	StudioID  uuid.UUID `json:"studio_id"`
}

func (c *CreateUserCommandResponse) FromEntity(u *entity.User) {
	c.ID = u.ID
	c.Name = u.FullName
	c.Email = u.Email
	c.Nickname = u.Nickname
	c.Type, _ = u.Type.Value()
	c.CreatedAt = u.CreatedAt
	c.StudioID = u.StudioID
}
