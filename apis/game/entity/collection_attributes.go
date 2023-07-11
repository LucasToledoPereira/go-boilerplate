package entity

import (
	"time"

	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CollectionAttributes struct {
	ID           uuid.UUID               `gorm:"primaryKey; index; unique; type:uuid;"`
	Name         string                  `gorm:"not null; index;"`
	Type         datatypes.AttributeType `gorm:"not null; type: attribute_type; default: STRING"`
	Value        string
	Collection   Collection
	CollectionID uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	models.Agents
}

func (CollectionAttributes) TableName() string { return "collection_attributes" }

func (g *CollectionAttributes) AfterUpdate(db *gorm.DB) (err error) {
	return gbp.SaveAuditRecord(g.UpdatedBy, g.TableName(), "UPDATE", g.ID, g, db)
}

func (g *CollectionAttributes) AfterCreate(db *gorm.DB) (err error) {
	return gbp.SaveAuditRecord(g.CreatedBy, g.TableName(), "INSERT", g.ID, g, db)
}

func (g *CollectionAttributes) AfterDelete(db *gorm.DB) (err error) {
	return gbp.SaveAuditRecord(g.DeletedBy, g.TableName(), "DELETE", g.ID, g, db)
}
