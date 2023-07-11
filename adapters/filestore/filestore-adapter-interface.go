package filestore

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type IFilestoreAdapter interface {
	New() error
	GetDomain() string
	Delete(key string, recordID uuid.UUID) error
	Save(file multipart.File, filename string, path string, recordID uuid.UUID) (string, error)
	GetTemporaryURL(key string, recordID uuid.UUID) (string, error)
	SaveAndGetTemporaryURL(file multipart.File, filename string, path string, recordID uuid.UUID) (string, string, error)
}
