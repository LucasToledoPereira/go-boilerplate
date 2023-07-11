package entity

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/audit"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Collection struct {
	ID               uuid.UUID `gorm:"primaryKey; index; unique; type:uuid;"`
	Name             string    `gorm:"not null; index;"`
	Symbol           string    `gorm:"not null;"`
	Description      string
	ShortDescription string
	CoverKey         string
	AvatarKey        string
	Address          string
	Attributes       []CollectionAttributes
	Game             Game
	GameID           uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	models.Agents
}

func (Collection) TableName() string { return "collections" }

func (g *Collection) AfterUpdate(db *gorm.DB) (err error) {
	return audit.Save(g.UpdatedBy, g.TableName(), "UPDATE", g.ID, g, db)
}

func (g *Collection) AfterCreate(db *gorm.DB) (err error) {
	return audit.Save(g.CreatedBy, g.TableName(), "INSERT", g.ID, g, db)
}

func (g *Collection) AfterDelete(db *gorm.DB) (err error) {
	return audit.Save(g.DeletedBy, g.TableName(), "DELETE", g.ID, g, db)
}
