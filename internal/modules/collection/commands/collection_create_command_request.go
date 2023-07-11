package collection_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateCollectionCommandRequest represents the request body for creating a new template
type CreateCollectionCommandRequest struct {
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
	Creator          uuid.UUID
	StudioID         uuid.UUID
	GameID           uuid.UUID
}

func (c *CreateCollectionCommandRequest) ToEntity() entity.Collection {
	tp := entity.Collection{
		ID:               uuid.New(),
		Name:             c.Name,
		Symbol:           c.Symbol,
		Description:      c.Description,
		ShortDescription: c.ShortDescription,
		GameID:           c.GameID,
	}
	tp.CreatedBy = c.Creator
	return tp
}

func (c *CreateCollectionCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}

func (c *CreateCollectionCommandRequest) Validate() (messages []string) {
	if c.Name == "" {
		messages = make([]string, 1)
		messages[0] = "name.required"
		return messages
	}

	if c.Symbol == "" {
		messages = make([]string, 1)
		messages[0] = "symbol.required"
		return messages
	}
	return nil
}
