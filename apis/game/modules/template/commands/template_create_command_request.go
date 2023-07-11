package template_commands

import (
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateTemplateRequest represents the request body for creating a new template
type CreateTemplateCommandRequest struct {
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Description  string `json:"description"`
	AnimationURL string `json:"animation_url"`
	ExternalURL  string `json:"external_url"`
	Category     string `json:"category"`
	Supply       string `json:"supply"`
	Creator      uuid.UUID
	StudioID     uuid.UUID
	GameID       uuid.UUID
}

func (c *CreateTemplateCommandRequest) ToEntity() entity.Template {
	tp := entity.Template{
		ID:           uuid.New(),
		Name:         c.Name,
		Symbol:       c.Symbol,
		Description:  c.Description,
		AnimationURL: c.AnimationURL,
		ExternalURL:  c.ExternalURL,
		Category:     c.Category,
		Supply:       c.Supply,
		GameID:       c.GameID,
	}
	tp.CreatedBy = c.Creator
	return tp
}

func (c *CreateTemplateCommandRequest) Bind(ctx *gin.Context) (err error) {
	return ctx.ShouldBindJSON(&c)
}

func (c *CreateTemplateCommandRequest) Validate() (messages []string) {
	if c.Name == "" {
		messages = make([]string, 1)
		messages[0] = "name.required"
		return messages
	}
	return nil
}
