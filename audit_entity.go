package gbp

import (
	"encoding/json"
	"time"

	"github.com/LucasToledoPereira/go-boilerplate/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Audit struct {
	ID        uuid.UUID `gorm:"primaryKey; index; unique; type:uuid;"`
	Table     string    `gorm:"not null;"`
	Action    string
	RecordID  uuid.UUID
	UserID    uuid.UUID
	Data      string
	CreatedAt time.Time
}

func (Audit) TableName() string { return "audit" }

func SaveAuditRecord(agentID uuid.UUID, table string, action string, record uuid.UUID, data interface{}, db *gorm.DB) error {
	if config.C.Settings.Audit {
		ad := &Audit{
			ID:       uuid.New(),
			Table:    table,
			Action:   action,
			RecordID: record,
			UserID:   agentID,
		}
		ad.setData(data)
		return db.Create(&ad).Error
	}
	return nil
}

func (a *Audit) setData(data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		a.Data = ""
	} else {
		a.Data = string(b)
	}
}
