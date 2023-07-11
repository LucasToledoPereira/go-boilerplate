package collection_handler

import (
	"errors"

	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/enums/codes"
	collection_commands "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/collection/commands"
	commands "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/collection/commands"
	interfaces "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/collection/interfaces"
	"github.com/LucasToledoPereira/go-boilerplate/internal/utils"
	"github.com/google/uuid"
)

type CollectionHandler struct {
	repository interfaces.ICollectionRepository
	filestore  gbp.IFilestoreAdapter
}

func NewCollectionHandler(repo interfaces.ICollectionRepository, filestore gbp.IFilestoreAdapter) (handler CollectionHandler) {
	return CollectionHandler{
		repository: repo,
		filestore:  filestore,
	}
}

func (h *CollectionHandler) Create(command commands.CreateCollectionCommandRequest) (res commands.CreateCollectionCommandResponse, err error) {
	belongs, err := h.repository.GameBelongsToStudio(command.GameID, command.StudioID)
	if err != nil {
		return res, err
	} else if !belongs {
		return res, errors.New(codes.GameNotBelongsToStudio.String())
	}

	collection := command.ToEntity()
	err = h.repository.Create(&collection)
	res.FromEntity(&collection)
	return res, err
}

func (h *CollectionHandler) List(gameId, studioId uuid.UUID) (res []commands.ListCollectionCommandResponse, err error) {
	belongs, err := h.repository.GameBelongsToStudio(gameId, studioId)
	if err != nil {
		return res, err
	} else if !belongs {
		return res, errors.New(codes.GameNotBelongsToStudio.String())
	}

	entities, err := h.repository.List(gameId)
	for _, attr := range entities {
		var lsc commands.ListCollectionCommandResponse
		lsc.FromEntity(attr)
		res = append(res, lsc)
	}
	return res, err
}

func (h *CollectionHandler) Read(collectionId, gameId, studioId uuid.UUID) (res commands.ReadCollectionCommandResponse, err error) {
	belongs, err := h.repository.GameBelongsToStudio(gameId, studioId)
	if err != nil {
		return res, err
	} else if !belongs {
		return res, errors.New(codes.GameNotBelongsToStudio.String())
	}

	collection, err := h.repository.Read(collectionId, gameId)

	coverKey := collection.CoverKey
	var coverurl string
	if coverKey != "" {
		coverurl, _ = h.filestore.GetTemporaryURL(coverKey, collection.ID)
	}

	avatarKey := collection.AvatarKey
	var avatarurl string
	if avatarKey != "" {
		avatarurl, _ = h.filestore.GetTemporaryURL(avatarKey, collection.ID)
	}

	res.FromEntity(&collection, coverurl, avatarurl)
	return res, err
}

func (h *CollectionHandler) Delete(collectionId, gameId, studioId, deleter uuid.UUID) (err error) {
	belongs, err := h.repository.GameBelongsToStudio(gameId, studioId)
	if err != nil {
		return err
	} else if !belongs {
		return errors.New(codes.GameNotBelongsToStudio.String())
	}

	err = h.repository.Delete(collectionId, gameId, deleter)
	return err
}

func (h *CollectionHandler) Update(command commands.UpdateCollectionCommandRequest) (res commands.UpdateCollectionCommandResponse, err error) {
	belongs, err := h.repository.GameBelongsToStudio(command.GameID, command.StudioID)
	if err != nil {
		return res, err
	} else if !belongs {
		return res, errors.New(codes.GameNotBelongsToStudio.String())
	}

	collection, err := h.repository.Read(command.CollectionID, command.GameID)
	if err != nil {
		return res, err
	}

	command.ToEntity(&collection)
	err = h.repository.Update(&collection)
	res.FromEntity(&collection)
	return res, err
}

func (h *CollectionHandler) UploadCover(command collection_commands.UploadFileCollectionCommandRequest) (res collection_commands.ReadCollectionCommandResponse, err error) {
	isImage := utils.IsImage(command.FileHeader.Filename)
	if !isImage {
		return res, errors.New(codes.FileNotImage.String())
	}

	belongs, err := h.repository.GameBelongsToStudio(command.GameID, command.StudioID)
	if err != nil {
		return res, err
	} else if !belongs {
		return res, errors.New(codes.GameNotBelongsToStudio.String())
	}

	collection, err := h.repository.Read(command.CollectionID, command.GameID)
	if err != nil {
		return res, err
	}

	oldKey := collection.CoverKey
	filekey, coverurl, _ := h.filestore.SaveAndGetTemporaryURL(command.File, command.FileHeader.Filename, collection.ID.String()+"/", collection.ID)
	err = h.repository.SetCoverKey(&collection, filekey)
	if oldKey != "" {
		_ = h.filestore.Delete(oldKey, collection.ID)
	}

	avatarKey := collection.AvatarKey
	var avatarurl string
	if avatarKey != "" {
		avatarurl, _ = h.filestore.GetTemporaryURL(avatarKey, collection.ID)
	}
	res.FromEntity(&collection, coverurl, avatarurl)
	return res, err
}

func (h *CollectionHandler) UploadAvatar(command collection_commands.UploadFileCollectionCommandRequest) (res collection_commands.ReadCollectionCommandResponse, err error) {
	isImage := utils.IsImage(command.FileHeader.Filename)
	if !isImage {
		return res, errors.New(codes.FileNotImage.String())
	}

	belongs, err := h.repository.GameBelongsToStudio(command.GameID, command.StudioID)
	if err != nil {
		return res, err
	} else if !belongs {
		return res, errors.New(codes.GameNotBelongsToStudio.String())
	}

	collection, err := h.repository.Read(command.CollectionID, command.GameID)
	if err != nil {
		return res, err
	}

	oldKey := collection.AvatarKey
	filekey, avatarurl, _ := h.filestore.SaveAndGetTemporaryURL(command.File, command.FileHeader.Filename, collection.ID.String()+"/", collection.ID)
	err = h.repository.SetAvatarKey(&collection, filekey)
	if oldKey != "" {
		_ = h.filestore.Delete(oldKey, collection.ID)
	}

	coverKey := collection.CoverKey
	var coverurl string
	if coverKey != "" {
		coverurl, _ = h.filestore.GetTemporaryURL(coverKey, collection.ID)
	}
	res.FromEntity(&collection, coverurl, avatarurl)
	return res, err
}

func (h *CollectionHandler) CreateAttribute(command commands.CreateCollectionAttributeCommandRequest) (res commands.CreateCollectionAttributeCommandResponse, err error) {
	belongs, err := h.repository.GameBelongsToStudio(command.GameID, command.StudioID)
	if err != nil {
		return res, err
	} else if !belongs {
		return res, errors.New(codes.GameNotBelongsToStudio.String())
	}

	belongs, err = h.repository.CollectionBelongsToGame(command.CollectionID, command.GameID)
	if err != nil {
		return res, err
	} else if !belongs {
		return res, errors.New(codes.CollectionNotBelongsToGame.String())
	}

	attribute := command.ToEntity()
	err = h.repository.CreateAttribute(&attribute)
	res.FromEntity(&attribute)
	return res, err
}

func (h *CollectionHandler) UpdateAttribute(command commands.UpdateCollectionAttributeCommandRequest) (res commands.UpdateCollectionAttributeCommandResponse, err error) {
	err = h.defaultChildrenChecks(command.CollectionID, command.GameID, command.StudioID)
	if err != nil {
		return res, err
	}

	attribute, err := h.repository.ReadAttribute(command.CollectionID, command.AttributeID)
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

func (h *CollectionHandler) ListAttributes(gameID, studioID, collectionID uuid.UUID) (res []commands.ListCollectionAttributesCommandResponse, err error) {
	err = h.defaultChildrenChecks(collectionID, gameID, studioID)
	if err != nil {
		return nil, err
	}

	attributes, err := h.repository.ListAttributes(collectionID)
	for _, attr := range attributes {
		var lsc commands.ListCollectionAttributesCommandResponse
		lsc.FromEntity(attr)
		res = append(res, lsc)
	}

	return res, err
}

func (h *CollectionHandler) ReadAttribute(gameID, studioID, collectionID, attributeID uuid.UUID) (res commands.ListCollectionAttributesCommandResponse, err error) {
	err = h.defaultChildrenChecks(collectionID, gameID, studioID)
	if err != nil {
		return res, err
	}
	attribute, err := h.repository.ReadAttribute(collectionID, attributeID)
	if err != nil {
		return res, err
	}

	res.FromEntity(attribute)

	return res, nil
}

func (h *CollectionHandler) DeleteAttribute(attributeID, collectionID, gameID, studioId, deleter uuid.UUID) error {
	err := h.defaultChildrenChecks(collectionID, gameID, studioId)
	if err != nil {
		return err
	}
	err = h.repository.DeleteAttribute(attributeID, collectionID, deleter)
	return err
}

func (h *CollectionHandler) defaultChildrenChecks(collectionID, gameID, studioID uuid.UUID) error {
	//Check if the game belongs to the same studio than the user
	belongs, err := h.repository.GameBelongsToStudio(gameID, studioID)
	if err != nil {
		return err
	} else if !belongs {
		return errors.New(codes.GameNotBelongsToStudio.String())
	}

	//Check if the template belongs to the game
	belongs, err = h.repository.CollectionBelongsToGame(collectionID, gameID)
	if err != nil {
		return err
	} else if !belongs {
		return errors.New(codes.CollectionNotBelongsToGame.String())
	}
	return nil
}
