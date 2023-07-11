package template_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateTemplateAttributeCommandRequest struct {
	Name        string                  `json:"name"`
	Type        datatypes.AttributeType `json:"type"`
	Value       string                  `json:"Value"`
	Updater     uuid.UUID
	StudioID    uuid.UUID
	GameID      uuid.UUID
	TemplateID  uuid.UUID
	AttributeID uuid.UUID
}

func (c *UpdateTemplateAttributeCommandRequest) ToEntity(tp *entity.TemplateAttributes) {
	if c.Name != "" {
		tp.Name = c.Name
	}
	if c.Type != "" {
		tp.Type = c.Type
	}
	tp.Value = c.Value
	tp.UpdatedBy = c.Updater
}

func (c *UpdateTemplateAttributeCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}
