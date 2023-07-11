package auth_controller

import (
	"errors"
	"net/http"

	codes "github.com/LucasToledoPereira/go-boilerplate/apis/game/enums/codes"
	auth_commands "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/auth/commands"
	handler "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/auth/handler"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	handler handler.AuthHandler
}

func NewAuthController(handler handler.AuthHandler) (c *AuthController) {
	return &AuthController{
		handler: handler,
	}
}

func (c AuthController) Register(ctx *gin.Context) {
	var commandRequest auth_commands.RegisterCommandRequest
	if err := commandRequest.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidRegisterFields.String(), false, []string{err.Error()}, nil))
		return
	}

	msgs := commandRequest.Validate()
	if msgs != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidRegisterFields.String(), false, msgs, nil))
		return
	}

	err := c.handler.Register(commandRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.StudioCreateFailed.String(), false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.StudioCreateSuccess.String(), true, nil, nil))
}

func (uc AuthController) Login(ctx *gin.Context) (interface{}, error) {
	var commandRequest auth_commands.LoginCommandRequest
	if err := commandRequest.Bind(ctx); err != nil {
		return nil, err
	}

	msgs := commandRequest.Validate()
	if msgs != nil {
		return nil, errors.New(codes.InvalidLoginFields.String())
	}

	resp, err := uc.handler.Login(commandRequest)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
