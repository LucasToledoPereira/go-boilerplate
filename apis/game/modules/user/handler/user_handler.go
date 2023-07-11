package user_handler

import (
	"errors"
	"mime/multipart"

	gbp "github.com/LucasToledoPereira/go-boilerplate"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/enums/codes"
	user_commands "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/user/commands"
	interfaces "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/user/interfaces"
	"github.com/LucasToledoPereira/go-boilerplate/internal/utils"
	"github.com/google/uuid"
)

type UserHandler struct {
	repository interfaces.IUserRepository
	filestore  gbp.IFilestoreAdapter
}

func NewUserHandler(repo interfaces.IUserRepository, filestore gbp.IFilestoreAdapter) (handler UserHandler) {
	return UserHandler{
		repository: repo,
		filestore:  filestore,
	}
}

func (h *UserHandler) ReadByIdentity(identity string) (res user_commands.ReadUserCommandResponse, err error) {
	user, err := h.repository.ReadByIdentity(identity)
	url, _ := h.filestore.GetTemporaryURL(user.Filekey, user.ID)
	res.FromEntity(&user, url)
	return res, err
}

func (h *UserHandler) DeleteByIdentity(identity string, deleter uuid.UUID) (err error) {
	err = h.repository.DeleteByIdentity(identity, deleter)
	return err
}

func (h *UserHandler) UpdateSelf(identity string, command user_commands.UpdateUserCommandRequest) (res user_commands.UpdateUserCommandResponse, err error) {

	user, err := h.repository.ReadByIdentity(identity)
	if err != nil {
		return user_commands.UpdateUserCommandResponse{}, err
	}

	if user.Nickname != command.Nickname && command.Nickname != "" {
		hasNick, err := h.repository.HasUserWithNickname(command.Nickname)
		if hasNick || err != nil {
			return user_commands.UpdateUserCommandResponse{}, errors.New(codes.UserWithNicknameAlreadyExists.String())
		}
	}

	if user.Email != command.Email && command.Email != "" {
		hasEmail, err := h.repository.HasUserWithEmail(command.Email)
		if hasEmail || err != nil {
			return user_commands.UpdateUserCommandResponse{}, errors.New(codes.UserWithEmailAlreadyExists.String())
		}
	}

	command.ToEntity(&user)
	err = h.repository.Update(&user)
	url, _ := h.filestore.GetTemporaryURL(user.Filekey, user.ID)
	res.FromEntity(&user, url)
	return res, err
}

func (h *UserHandler) List(studioID uuid.UUID) (users []user_commands.ListUserCommandResponse, err error) {
	entities, err := h.repository.List(studioID)

	for _, user := range entities {
		var lsc user_commands.ListUserCommandResponse
		url, _ := h.filestore.GetTemporaryURL(user.Filekey, user.ID)
		lsc.FromEntity(user, url)
		users = append(users, lsc)
	}
	return users, err
}

func (h *UserHandler) Create(studioID uuid.UUID, command user_commands.CreateUserCommandRequest) (res user_commands.CreateUserCommandResponse, err error) {
	hasNick, err := h.repository.HasUserWithNickname(command.Nickname)
	if hasNick || err != nil {
		return user_commands.CreateUserCommandResponse{}, errors.New(codes.UserWithNicknameAlreadyExists.String())
	}

	hasEmail, err := h.repository.HasUserWithEmail(command.Email)
	if hasEmail || err != nil {
		return user_commands.CreateUserCommandResponse{}, errors.New(codes.UserWithEmailAlreadyExists.String())
	}

	user := command.ToEntity(studioID)
	err = h.repository.Create(&user)
	res.FromEntity(&user)
	return res, err
}

func (h *UserHandler) Delete(userID uuid.UUID, studioID uuid.UUID, deleter uuid.UUID) (err error) {
	err = h.repository.Delete(userID, studioID, deleter)
	return err
}

func (h *UserHandler) Update(id uuid.UUID, studioID uuid.UUID, command user_commands.UpdateUserCommandRequest) (res user_commands.UpdateUserCommandResponse, err error) {
	user, err := h.repository.Read(id, studioID)
	if err != nil {
		return user_commands.UpdateUserCommandResponse{}, err
	}

	if user.Nickname != command.Nickname && command.Nickname != "" {
		hasNick, err := h.repository.HasUserWithNickname(command.Nickname)
		if hasNick || err != nil {
			return user_commands.UpdateUserCommandResponse{}, errors.New(codes.UserWithNicknameAlreadyExists.String())
		}
	}

	if user.Email != command.Email && command.Email != "" {
		hasEmail, err := h.repository.HasUserWithEmail(command.Email)
		if hasEmail || err != nil {
			return user_commands.UpdateUserCommandResponse{}, errors.New(codes.UserWithEmailAlreadyExists.String())
		}
	}

	command.ToEntity(&user)
	err = h.repository.Update(&user)
	url, _ := h.filestore.GetTemporaryURL(user.Filekey, user.ID)
	res.FromEntity(&user, url)
	return res, err
}

func (h *UserHandler) Read(id uuid.UUID, studioID uuid.UUID) (res user_commands.ReadUserCommandResponse, err error) {
	user, err := h.repository.Read(id, studioID)
	url, _ := h.filestore.GetTemporaryURL(user.Filekey, user.ID)
	res.FromEntity(&user, url)
	return res, err
}

func (h *UserHandler) Upload(id uuid.UUID, studioID uuid.UUID, file multipart.File, fileHandler *multipart.FileHeader) (res user_commands.ReadUserCommandResponse, err error) {
	isImage := utils.IsImage(fileHandler.Filename)
	if !isImage {
		return res, errors.New(codes.FileNotImage.String())
	}

	user, err := h.repository.Read(id, studioID)
	if err != nil {
		return res, err
	}

	oldKey := user.Filekey
	filekey, url, _ := h.filestore.SaveAndGetTemporaryURL(file, fileHandler.Filename, user.ID.String()+"/", user.ID)
	err = h.repository.UpdateImage(user, filekey)
	if oldKey != "" {
		_ = h.filestore.Delete(oldKey, user.ID)
	}
	res.FromEntity(&user, url)
	return res, err
}

func (h *UserHandler) UploadSelf(identity string, file multipart.File, fileHandler *multipart.FileHeader) (res user_commands.ReadUserCommandResponse, err error) {
	isImage := utils.IsImage(fileHandler.Filename)
	if !isImage {
		return res, errors.New(codes.FileNotImage.String())
	}

	user, err := h.repository.ReadByIdentity(identity)
	if err != nil {
		return res, err
	}

	oldKey := user.Filekey
	filekey, url, _ := h.filestore.SaveAndGetTemporaryURL(file, fileHandler.Filename, user.ID.String()+"/", user.ID)
	err = h.repository.UpdateImage(user, filekey)
	if oldKey != "" {
		_ = h.filestore.Delete(oldKey, user.ID)
	}
	res.FromEntity(&user, url)
	return res, err
}
