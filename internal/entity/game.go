package entity

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/audit"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Game struct {
	ID          uuid.UUID `gorm:"primaryKey; index; unique; type:uuid;"`
	Title       string    `gorm:"not null;"`
	Description string
	Website     string
	StudioID    uuid.UUID `gorm:"type:uuid;not null"`
	Studio      Studio
	Templates   []Template
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	models.Agents
}

func (Game) TableName() string { return "games" }

func (g *Game) AfterUpdate(db *gorm.DB) (err error) {
	return audit.Save(g.UpdatedBy, g.TableName(), "UPDATE", g.ID, g, db)
}

func (g *Game) AfterCreate(db *gorm.DB) (err error) {
	return audit.Save(g.CreatedBy, g.TableName(), "INSERT", g.ID, g, db)
}

func (g *Game) AfterDelete(db *gorm.DB) (err error) {
	return audit.Save(g.DeletedBy, g.TableName(), "DELETE", g.ID, g, db)
}
