package collection_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateCollectionAttributeCommandRequest struct {
	Name         string                  `json:"name"`
	Type         datatypes.AttributeType `json:"type"`
	Value        string                  `json:"value"`
	Creator      uuid.UUID
	StudioID     uuid.UUID
	GameID       uuid.UUID
	CollectionID uuid.UUID
}

func (c *CreateCollectionAttributeCommandRequest) ToEntity() entity.CollectionAttributes {
	tp := entity.CollectionAttributes{
		ID:           uuid.New(),
		Name:         c.Name,
		Type:         c.Type,
		Value:        c.Value,
		CollectionID: c.CollectionID,
	}
	tp.CreatedBy = c.Creator
	return tp
}

func (c *CreateCollectionAttributeCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}

func (c *CreateCollectionAttributeCommandRequest) Validate() (messages []string) {
	if c.Name == "" {
		messages = make([]string, 1)
		messages[0] = "name.required"
		return messages
	}
	if c.Type == "" {
		messages = make([]string, 1)
		messages[0] = "type.required"
		return messages
	}
	if c.Value == "" {
		messages = make([]string, 1)
		messages[0] = "value.required"
		return messages
	}
	//TODO is missing the type validation, like, if the type is NUMBER, we must check if the value sent is a valid number
	return nil
}
