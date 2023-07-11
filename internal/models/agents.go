package models

import "github.com/google/uuid"

type Agents struct {
	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	DeletedBy uuid.UUID
}
