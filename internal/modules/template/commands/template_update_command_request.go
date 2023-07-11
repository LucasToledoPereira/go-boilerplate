package template_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateTemplateCommandRequest struct {
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Description  string `json:"description"`
	AnimationURL string `json:"animation_url"`
	ExternalURL  string `json:"external_url"`
	Category     string `json:"category"`
	Supply       string `json:"supply"`
	Updater      uuid.UUID
	StudioID     uuid.UUID
	GameID       uuid.UUID
	TemplateID   uuid.UUID
}

func (c *UpdateTemplateCommandRequest) ToEntity(tp *entity.Template) {
	if c.Name != "" {
		tp.Name = c.Name
	}
	tp.Symbol = c.Symbol
	tp.Supply = c.Supply
	tp.Description = c.Description
	tp.AnimationURL = c.AnimationURL
	tp.ExternalURL = c.ExternalURL
	tp.Category = c.Category
	tp.UpdatedBy = c.Updater
}

func (c *UpdateTemplateCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}
