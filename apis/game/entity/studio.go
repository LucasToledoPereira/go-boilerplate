package entity

import (
	"time"

	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Studio struct {
	ID          uuid.UUID `gorm:"primaryKey; index; unique; type:uuid;"`
	Name        string    `gorm:"not null; index; unique;"`
	Description string
	Website     string
	Instagram   string
	Filekey     string
	Games       []Game
	Users       []User
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	models.Agents
}

func (Studio) TableName() string { return "studios" }

func (st *Studio) AfterUpdate(db *gorm.DB) (err error) {
	return gbp.SaveAuditRecord(st.UpdatedBy, st.TableName(), "UPDATE", st.ID, st, db)
}

func (st *Studio) AfterCreate(db *gorm.DB) (err error) {
	return gbp.SaveAuditRecord(st.CreatedBy, st.TableName(), "INSERT", st.ID, st, db)
}

func (st *Studio) AfterDelete(db *gorm.DB) (err error) {
	return gbp.SaveAuditRecord(st.DeletedBy, st.TableName(), "DELETE", st.ID, st, db)
}
