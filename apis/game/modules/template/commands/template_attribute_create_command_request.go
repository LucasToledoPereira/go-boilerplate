package template_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateTemplateAttributeCommandRequest struct {
	Name       string                  `json:"name"`
	Type       datatypes.AttributeType `json:"type"`
	Value      string                  `json:"value"`
	Creator    uuid.UUID
	StudioID   uuid.UUID
	GameID     uuid.UUID
	TemplateID uuid.UUID
}

func (c *CreateTemplateAttributeCommandRequest) ToEntity() entity.TemplateAttributes {
	tp := entity.TemplateAttributes{
		ID:         uuid.New(),
		Name:       c.Name,
		Type:       c.Type,
		Value:      c.Value,
		TemplateID: c.TemplateID,
	}
	tp.CreatedBy = c.Creator
	return tp
}

func (c *CreateTemplateAttributeCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}

func (c *CreateTemplateAttributeCommandRequest) Validate() (messages []string) {
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
