package studio_handler

import (
	"errors"
	"mime/multipart"

	"github.com/google/uuid"

	"github.com/LucasToledoPereira/go-boilerplate/adapters/filestore"
	"github.com/LucasToledoPereira/go-boilerplate/internal/enums/codes"
	studio_commands "github.com/LucasToledoPereira/go-boilerplate/internal/modules/studio/commands"
	interfaces "github.com/LucasToledoPereira/go-boilerplate/internal/modules/studio/interfaces"
	"github.com/LucasToledoPereira/go-boilerplate/internal/utils"
)

type StudioHandler struct {
	repository interfaces.IStudioRepository
	filestore  filestore.IFilestoreAdapter
}

func NewStudioHandler(repo interfaces.IStudioRepository, filestore filestore.IFilestoreAdapter) (handler StudioHandler) {
	return StudioHandler{
		repository: repo,
		filestore:  filestore,
	}
}

func (sh *StudioHandler) Create(command studio_commands.CreateStudioCommandRequest) (res studio_commands.CreateStudioCommandResponse, err error) {
	studio := command.ToEntity()
	err = sh.repository.Create(studio)
	res.FromEntity(studio)
	return res, err
}

func (sh *StudioHandler) Update(id uuid.UUID, command studio_commands.UpdateStudioCommandRequest) (res studio_commands.UpdateStudioCommandResponse, err error) {
	studio, err := sh.repository.Read(id)
	if err != nil {
		return studio_commands.UpdateStudioCommandResponse{}, err
	}

	command.ToEntity(&studio)
	err = sh.repository.Update(&studio)
	res.FromEntity(&studio)
	return res, err
}

func (sh *StudioHandler) Delete(id uuid.UUID) error {
	err := sh.repository.Delete(id)
	return err
}

func (sh *StudioHandler) List() (studios []studio_commands.ListStudioCommandResponse, err error) {
	entities, err := sh.repository.List()
	studios = studio_commands.ListResponseFromEntities(entities)
	return studios, err
}

func (sh *StudioHandler) Read(id uuid.UUID) (res studio_commands.ReadStudioCommandResponse, err error) {
	studio, err := sh.repository.Read(id)
	url, _ := sh.filestore.GetTemporaryURL(studio.Filekey, studio.ID)
	res.FromEntity(&studio, url)
	return res, err
}

func (sh *StudioHandler) Upload(studioID uuid.UUID, file multipart.File, fileHandler *multipart.FileHeader) (res studio_commands.ReadStudioCommandResponse, err error) {
	isImage := utils.IsImage(fileHandler.Filename)
	if !isImage {
		return res, errors.New(codes.FileNotImage.String())
	}

	studio, err := sh.repository.Read(studioID)
	if err != nil {
		return res, err
	}

	oldKey := studio.Filekey
	filekey, url, _ := sh.filestore.SaveAndGetTemporaryURL(file, fileHandler.Filename, studio.ID.String()+"/", studio.ID)
	err = sh.repository.UpdateImage(studio, filekey)
	if oldKey != "" {
		_ = sh.filestore.Delete(oldKey, studio.ID)
	}
	res.FromEntity(&studio, url)
	return res, err
}
