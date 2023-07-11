package collection_commands

import (
	"mime/multipart"

	"github.com/google/uuid"
)

// command Request to create a user
type UploadFileCollectionCommandRequest struct {
	ID           uuid.UUID
	GameID       uuid.UUID
	StudioID     uuid.UUID
	CollectionID uuid.UUID
	File         multipart.File
	FileHeader   *multipart.FileHeader
}
