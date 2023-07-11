package game_commands

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/google/uuid"
)

type ListGameCommandResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   uuid.UUID `json:"created_by"`
}

func (c *ListGameCommandResponse) FromEntity(g entity.Game) {
	c.ID = g.ID
	c.Title = g.Title
	c.Description = g.Description
	c.Website = g.Website
	c.CreatedAt = g.CreatedAt
	c.UpdatedAt = g.UpdatedAt
	c.CreatedBy = g.CreatedBy
}

func ListResponseFromEntities(games []entity.Game) (res []ListGameCommandResponse) {
	for _, game := range games {
		var lsc ListGameCommandResponse
		lsc.FromEntity(game)
		res = append(res, lsc)
	}
	return res
}
