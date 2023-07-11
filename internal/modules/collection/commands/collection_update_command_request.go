package collection_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateCollectionCommandRequest struct {
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
	Updater          uuid.UUID
	StudioID         uuid.UUID
	GameID           uuid.UUID
	CollectionID     uuid.UUID
}

func (c *UpdateCollectionCommandRequest) ToEntity(tp *entity.Collection) {
	if c.Name != "" {
		tp.Name = c.Name
	}
	tp.Symbol = c.Symbol
	tp.Description = c.Description
	tp.ShortDescription = c.ShortDescription
	tp.UpdatedBy = c.Updater
}

func (c *UpdateCollectionCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}
