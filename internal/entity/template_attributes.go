package entity

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/audit"
	"github.com/LucasToledoPereira/go-boilerplate/internal/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TemplateAttributes struct {
	ID         uuid.UUID               `gorm:"primaryKey; index; unique; type:uuid;"`
	Name       string                  `gorm:"not null; index;"`
	Type       datatypes.AttributeType `gorm:"not null; type: attribute_type; default: STRING"`
	Value      string
	Template   Template
	TemplateID uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	models.Agents
}

func (TemplateAttributes) TableName() string { return "template_attributes" }

func (g *TemplateAttributes) AfterUpdate(db *gorm.DB) (err error) {
	return audit.Save(g.UpdatedBy, g.TableName(), "UPDATE", g.ID, g, db)
}

func (g *TemplateAttributes) AfterCreate(db *gorm.DB) (err error) {
	return audit.Save(g.CreatedBy, g.TableName(), "INSERT", g.ID, g, db)
}

func (g *TemplateAttributes) AfterDelete(db *gorm.DB) (err error) {
	return audit.Save(g.DeletedBy, g.TableName(), "DELETE", g.ID, g, db)
}
