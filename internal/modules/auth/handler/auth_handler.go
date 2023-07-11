package auth_handler

import (
	"errors"

	"github.com/LucasToledoPereira/go-boilerplate/adapters/filestore"
	"github.com/LucasToledoPereira/go-boilerplate/internal/enums/codes"
	auth_commands "github.com/LucasToledoPereira/go-boilerplate/internal/modules/auth/commands"
	interfaces "github.com/LucasToledoPereira/go-boilerplate/internal/modules/auth/interfaces"
)

type AuthHandler struct {
	repository interfaces.IAuthRepository
	filestore  filestore.IFilestoreAdapter
}

func NewAuthHandler(repo interfaces.IAuthRepository, filestore filestore.IFilestoreAdapter) (handler AuthHandler) {
	return AuthHandler{
		repository: repo,
		filestore:  filestore,
	}
}

func (sh AuthHandler) Register(command auth_commands.RegisterCommandRequest) (err error) {
	//Verify if user already exists with e-mail or nickname, if true, return error
	userexists, err := sh.repository.HasUserWithEmailOrNickname(command.Email, command.Nickname)
	if userexists {
		return errors.New(codes.UserAlreadyExists.String())
	} else if err != nil {
		return err
	}

	//Verify if studio already exists with name, if true, return error
	studioexists, err := sh.repository.HasStudioWithName(command.StudioName)
	if studioexists {
		return errors.New(codes.StudioAlreadyExists.String())
	} else if err != nil {
		return err
	}

	user := command.ToEntity()
	err = sh.repository.Register(user)
	return err
}

func (sh AuthHandler) Login(command auth_commands.LoginCommandRequest) (response auth_commands.LoginCommandResponse, err error) {
	user, err := sh.repository.ReadByEmailOrNickname(command.Nickname)

	if err != nil {
		return response, err
	}

	pwdmatch := command.CheckPassword(user.Password)
	if pwdmatch {
		url, _ := sh.filestore.GetTemporaryURL(user.Filekey, user.ID)
		response.FromEntity(&user, url)
		return response, nil
	}

	return response, errors.New(codes.WrongPassword.String())
}
