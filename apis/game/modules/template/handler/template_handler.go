package temaplte_handler

import (
	"errors"

	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/entity"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/enums/codes"
	commands "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/template/commands"
	interfaces "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/template/interfaces"
	"github.com/LucasToledoPereira/go-boilerplate/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TemplateHandler struct {
	repository interfaces.ITemplateRepository
	filestore  gbp.IFilestoreAdapter
}

func NewTemplateHandler(repo interfaces.ITemplateRepository, filestore gbp.IFilestoreAdapter) (handler TemplateHandler) {
	return TemplateHandler{
		repository: repo,
		filestore:  filestore,
	}
}

func (h *TemplateHandler) Create(command commands.CreateTemplateCommandRequest) (res commands.CreateTemplateCommandResponse, err error) {
	belongs, err := h.repository.GameBelongsToStudio(command.GameID, command.StudioID)
	if err != nil {
		return res, err
	} else if !belongs {
		return res, errors.New(codes.GameNotBelongsToStudio.String())
	}

	template := command.ToEntity()
	err = h.repository.Create(&template)
	res.FromEntity(&template)
	return res, err
}

func (h *TemplateHandler) List(gameId uuid.UUID, studioId uuid.UUID) (templates []commands.ListTemplateCommandResponse, err error) {
	belongs, err := h.repository.GameBelongsToStudio(gameId, studioId)
	if err != nil {
		return templates, err
	} else if !belongs {
		return templates, errors.New(codes.GameNotBelongsToStudio.String())
	}

	entities, err := h.repository.List(gameId)
	templates = commands.ListResponseFromEntities(entities)
	return templates, err
}

func (h *TemplateHandler) Read(templateId uuid.UUID, gameId uuid.UUID, studioId uuid.UUID) (res commands.ReadTemplateCommandResponse, err error) {
	belongs, err := h.repository.GameBelongsToStudio(gameId, studioId)
	if err != nil {
		return res, err
	} else if !belongs {
		return res, errors.New(codes.GameNotBelongsToStudio.String())
	}
	game, err := h.repository.Read(templateId, gameId)
	res.FromEntity(&game)
	return res, err
}

func (h *TemplateHandler) Update(command commands.UpdateTemplateCommandRequest) (res commands.UpdateTemplateCommandResponse, err error) {
	belongs, err := h.repository.GameBelongsToStudio(command.GameID, command.StudioID)
	if err != nil {
		return res, err
	} else if !belongs {
		return res, errors.New(codes.GameNotBelongsToStudio.String())
	}

	template, err := h.repository.Read(command.TemplateID, command.GameID)
	if err != nil {
		return res, err
	}

	command.ToEntity(&template)
	err = h.repository.Update(&template)
	res.FromEntity(&template)
	return res, err
}

func (h *TemplateHandler) Delete(templateId uuid.UUID, gameId uuid.UUID, studioId uuid.UUID, deleter uuid.UUID) (err error) {
	belongs, err := h.repository.GameBelongsToStudio(gameId, studioId)
	if err != nil {
		return err
	} else if !belongs {
		return errors.New(codes.GameNotBelongsToStudio.String())
	}

	err = h.repository.Delete(templateId, gameId, deleter)
	return err
}

func (h *TemplateHandler) CreateAttribute(command commands.CreateTemplateAttributeCommandRequest) (res commands.CreateTemplateAttributeCommandResponse, err error) {
	belongs, err := h.repository.GameBelongsToStudio(command.GameID, command.StudioID)
	if err != nil {
		return res, err
	} else if !belongs {
		return res, errors.New(codes.GameNotBelongsToStudio.String())
	}

	belongs, err = h.repository.TemplateBelongsToGame(command.TemplateID, command.GameID)
	if err != nil {
		return res, err
	} else if !belongs {
		return res, errors.New(codes.TemplateNotBelongsToGame.String())
	}

	attribute := command.ToEntity()
	err = h.repository.CreateAttribute(&attribute)
	res.FromEntity(&attribute)
	return res, err
}

func (h *TemplateHandler) ListAttributes(gameID, studioID, templateID uuid.UUID) (res []commands.ListTemplateAttributeCommandResponse, err error) {
	err = h.defaultChildrenChecks(templateID, gameID, studioID)
	if err != nil {
		return nil, err
	}

	attributes, err := h.repository.ListAttributes(templateID)
	for _, attr := range attributes {
		var lsc commands.ListTemplateAttributeCommandResponse
		lsc.FromEntity(attr)
		res = append(res, lsc)
	}

	return res, err
}

func (h *TemplateHandler) ReadAttribute(gameID, studioID, templateID, attributeID uuid.UUID) (res commands.ListTemplateAttributeCommandResponse, err error) {
	err = h.defaultChildrenChecks(templateID, gameID, studioID)
	if err != nil {
		return res, err
	}
	attribute, err := h.repository.ReadAttribute(templateID, attributeID)
	if err != nil {
		return res, err
	}

	res.FromEntity(attribute)

	return res, nil
}

func (h *TemplateHandler) UpdateAttribute(command commands.UpdateTemplateAttributeCommandRequest) (res commands.UpdateTemplateAttributeCommandResponse, err error) {
	err = h.defaultChildrenChecks(command.TemplateID, command.GameID, command.StudioID)
	if err != nil {
		return res, err
	}

	attribute, err := h.repository.ReadAttribute(command.TemplateID, command.AttributeID)
	if err != nil {
		return res, err
	}

	command.ToEntity(&attribute)
	err = h.repository.UpdateAttribute(&attribute)
	if err != nil {
		return res, err
	}
	res.FromEntity(&attribute)
	return res, nil

}

func (h *TemplateHandler) DeleteAttribute(attributeID, templateID, gameID, studioId, deleter uuid.UUID) error {
	err := h.defaultChildrenChecks(templateID, gameID, studioId)
	if err != nil {
		return err
	}
	err = h.repository.DeleteAttribute(attributeID, templateID, deleter)
	return err
}

func (h *TemplateHandler) Upload(command commands.UploadFileTemplateCommandRequest) (res commands.UploadFileTemplateCommandResponse, err error) {
	err = h.defaultChildrenChecks(command.TemplateID, command.GameID, command.StudioID)
	if err != nil {
		return res, err
	}

	var oldKey string = ""
	var file entity.TemplateFiles
	var notFound bool = true

	if command.Type.IsType(datatypes.COVER) {
		isImage := utils.IsImage(command.FileHeader.Filename)
		if !isImage {
			return res, errors.New(codes.FileNotImage.String())
		}

		file, err = h.repository.FirstFileByType(command.TemplateID, command.Type)
		if err != nil {
			notFound = errors.Is(err, gorm.ErrRecordNotFound)

			if !notFound {
				return res, err
			}
		}

		oldKey = file.Key
	}

	if notFound {
		file = command.ToNewEntity()
	} else {
		command.ToEntity(&file)
	}
	filekey, url, _ := h.filestore.SaveAndGetTemporaryURL(command.File, command.FileHeader.Filename, file.TemplateID.String()+"/"+file.ID.String()+"/", file.ID)
	file.Key = filekey
	err = h.repository.SaveFile(&file)
	if err != nil {
		return res, err
	}

	if oldKey != "" {
		_ = h.filestore.Delete(oldKey, file.ID)
	}

	res.FromEntity(file, url)
	return res, nil
}

func (h *TemplateHandler) ListFilesIgnoringCover(gameID, studioID, templateID uuid.UUID) (res []commands.UploadFileTemplateCommandResponse, err error) {
	err = h.defaultChildrenChecks(templateID, gameID, studioID)
	if err != nil {
		return nil, err
	}

	files, err := h.repository.ListFilesIgnoringType(templateID, datatypes.COVER)
	for _, file := range files {
		var lsc commands.UploadFileTemplateCommandResponse
		url, _ := h.filestore.GetTemporaryURL(file.Key, file.ID)
		lsc.FromEntity(file, url)
		res = append(res, lsc)
	}

	return res, err
}

func (h *TemplateHandler) DeleteFile(fileID, templateID, gameID, studioID uuid.UUID) error {
	err := h.defaultChildrenChecks(templateID, gameID, studioID)
	if err != nil {
		return err
	}
	file, err := h.repository.GetFile(fileID, templateID)
	if err != nil {
		return err
	}

	err = h.repository.DeleteFile(fileID, templateID)
	if err != nil {
		return err
	}

	if file.Key != "" {
		_ = h.filestore.Delete(file.Key, file.ID)
	}
	return nil
}

func (h *TemplateHandler) defaultChildrenChecks(templateID, gameID, studioID uuid.UUID) error {
	//Check if the game belongs to the same studio than the user
	belongs, err := h.repository.GameBelongsToStudio(gameID, studioID)
	if err != nil {
		return err
	} else if !belongs {
		return errors.New(codes.GameNotBelongsToStudio.String())
	}

	//Check if the template belongs to the game
	belongs, err = h.repository.TemplateBelongsToGame(templateID, gameID)
	if err != nil {
		return err
	} else if !belongs {
		return errors.New(codes.TemplateNotBelongsToGame.String())
	}
	return nil
}
