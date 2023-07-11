package entity

import (
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/internal/audit"
	"github.com/LucasToledoPereira/go-boilerplate/internal/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID          `gorm:"primaryKey; index; unique; type:uuid;"`
	Email     string             `gorm:"not null; unique; index"`
	Nickname  string             `gorm:"not null; unique;"`
	FullName  string             `gorm:"not null;"`
	Password  string             `gorm:"not null;"`
	Type      datatypes.UserRole `gorm:"not null; type: user_role; default: COMMON"`
	Filekey   string
	StudioID  uuid.UUID `gorm:"type:uuid;not null"`
	Studio    Studio
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	models.Agents
}

func (User) TableName() string { return "users" }

func (u *User) AfterUpdate(db *gorm.DB) (err error) {
	return audit.Save(u.UpdatedBy, u.TableName(), "UPDATE", u.ID, u, db)
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
	return audit.Save(u.CreatedBy, u.TableName(), "INSERT", u.ID, u, db)
}

func (u *User) AfterDelete(db *gorm.DB) (err error) {
	return audit.Save(u.DeletedBy, u.TableName(), "DELETE", u.ID, u, db)
}
