package entity

import (
	"time"

	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TemplateFiles struct {
	ID         uuid.UUID          `gorm:"primaryKey; index; unique; type:uuid;"`
	Key        string             `gorm:"not null; index;"`
	Type       datatypes.FileType `gorm:"not null; type: file_type; default: ANY"`
	Template   Template
	TemplateID uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	models.Agents
}

func (TemplateFiles) TableName() string { return "template_files" }

func (g *TemplateFiles) AfterUpdate(db *gorm.DB) (err error) {
	return gbp.SaveAuditRecord(g.UpdatedBy, g.TableName(), "UPDATE", g.ID, g, db)
}

func (g *TemplateFiles) AfterCreate(db *gorm.DB) (err error) {
	return gbp.SaveAuditRecord(g.CreatedBy, g.TableName(), "INSERT", g.ID, g, db)
}

func (g *TemplateFiles) AfterDelete(db *gorm.DB) (err error) {
	return gbp.SaveAuditRecord(g.DeletedBy, g.TableName(), "DELETE", g.ID, g, db)
}
