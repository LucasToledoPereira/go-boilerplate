package template_interfaces

import (
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/google/uuid"
)

type ITemplateRepository interface {
	Create(template *entity.Template) error
	Update(template *entity.Template) error
	Delete(template uuid.UUID, gameId uuid.UUID, deleter uuid.UUID) error
	List(studioID uuid.UUID) (templates []entity.Template, err error)
	Read(id uuid.UUID, studioID uuid.UUID) (template entity.Template, err error)
	GameBelongsToStudio(gameID uuid.UUID, studioID uuid.UUID) (bool, error)
	TemplateBelongsToGame(templateID uuid.UUID, gameID uuid.UUID) (bool, error)
	CreateAttribute(attribute *entity.TemplateAttributes) error
	UpdateAttribute(attribute *entity.TemplateAttributes) error
	ListAttributes(templateID uuid.UUID) (attributes []entity.TemplateAttributes, err error)
	ReadAttribute(templateID, attributeID uuid.UUID) (attribute entity.TemplateAttributes, err error)
	DeleteAttribute(attributeID, templateID, deleter uuid.UUID) error
	FirstFileByType(templateID uuid.UUID, filetype datatypes.FileType) (file entity.TemplateFiles, err error)
	SaveFile(file *entity.TemplateFiles) error
	ListFilesIgnoringType(templateID uuid.UUID, filetypeToIgnore datatypes.FileType) (files []entity.TemplateFiles, err error)
	GetFile(fileID, templateID uuid.UUID) (file entity.TemplateFiles, err error)
	DeleteFile(fileID, templateID uuid.UUID) error
}
