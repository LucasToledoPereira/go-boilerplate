package template_commands

import (
	"mime/multipart"

	"github.com/LucasToledoPereira/go-boilerplate/internal/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/internal/entity"
	"github.com/google/uuid"
)

// command Request to create a user
type UploadFileTemplateCommandRequest struct {
	ID         uuid.UUID
	GameID     uuid.UUID
	StudioID   uuid.UUID
	TemplateID uuid.UUID
	File       multipart.File
	FileHeader *multipart.FileHeader
	Type       datatypes.FileType
	Creator    uuid.UUID
}

func (c *UploadFileTemplateCommandRequest) ToEntity(tp *entity.TemplateFiles) {
	tp.UpdatedBy = c.Creator
}

func (c *UploadFileTemplateCommandRequest) ToNewEntity() entity.TemplateFiles {
	file := entity.TemplateFiles{
		ID:         uuid.New(),
		Type:       c.Type,
		TemplateID: c.TemplateID,
	}
	file.CreatedBy = c.Creator
	return file
}
