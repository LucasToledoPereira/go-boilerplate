package entity

import (
	"time"

	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Template struct {
	ID           uuid.UUID `gorm:"primaryKey; index; unique; type:uuid;"`
	Name         string    `gorm:"not null; index;"`
	Symbol       string
	Description  string
	AnimationURL string
	ExternalURL  string
	Category     string
	Supply       string
	Cover        string `gorm:"-"`
	Attributes   []TemplateAttributes
	Files        []TemplateFiles
	Game         Game
	GameID       uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	models.Agents
}

func (Template) TableName() string { return "templates" }

func (g *Template) AfterUpdate(db *gorm.DB) (err error) {
	return gbp.SaveAuditRecord(g.UpdatedBy, g.TableName(), "UPDATE", g.ID, g, db)
}

func (g *Template) AfterCreate(db *gorm.DB) (err error) {
	return gbp.SaveAuditRecord(g.CreatedBy, g.TableName(), "INSERT", g.ID, g, db)
}

func (g *Template) AfterDelete(db *gorm.DB) (err error) {
	return gbp.SaveAuditRecord(g.DeletedBy, g.TableName(), "DELETE", g.ID, g, db)
}
