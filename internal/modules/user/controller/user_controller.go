package user_controller

import (
	"errors"
	"net/http"

	"github.com/LucasToledoPereira/go-boilerplate/internal/enums/codes"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	user_commands "github.com/LucasToledoPereira/go-boilerplate/internal/modules/user/commands"
	handler "github.com/LucasToledoPereira/go-boilerplate/internal/modules/user/handler"
	"github.com/LucasToledoPereira/go-boilerplate/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserController struct {
	handler handler.UserHandler
}

func NewUserController(handler handler.UserHandler) (c UserController) {
	return UserController{
		handler: handler,
	}
}

func (sc *UserController) Create(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		id, err := utils.GetUserStudioID(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
			return
		}

		var commandRequest user_commands.CreateUserCommandRequest
		if err := commandRequest.Bind(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserInvalidFields, false, []string{err.Error()}, nil))
			return
		}

		creatorID, _ := utils.GetUserID(ctx)
		commandRequest.Creator = creatorID
		msgs := commandRequest.Validate()
		if msgs != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserInvalidFields, false, msgs, nil))
			return
		}

		res, err := sc.handler.Create(id, commandRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserCreateFailed, false, []string{err.Error()}, nil))
			return
		}

		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.UserCreateSuccess, true, nil, res))
		return
	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
}

func (sc *UserController) List(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		studioID, err := utils.GetUserStudioID(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.NewResultWrapper[any](codes.InvalidParamID, false, []string{err.Error()}, nil))
			return
		}
		res, err := sc.handler.List(studioID)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.NewResultWrapper[any](codes.UserListsFailed, false, []string{err.Error()}, nil))
			return
		}
		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.UserListsSuccess, true, nil, res))
		return
	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
}

func (sc *UserController) ReadInfo(ctx *gin.Context) {
	id := utils.GetUserIdentity(ctx)
	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserInvalidIdentity, false, nil, nil))
		return
	}

	res, err := sc.handler.ReadByIdentity(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserNotFound, false, nil, nil))
			return
		}
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.UserReadSuccess, true, nil, res))
}

func (sc *UserController) DeleteSelf(ctx *gin.Context) {
	id := utils.GetUserIdentity(ctx)
	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserInvalidIdentity, false, nil, nil))
		return
	}

	deleter, _ := utils.GetUserID(ctx)
	err := sc.handler.DeleteByIdentity(id, deleter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.UserDeleteSuccess, true, nil, nil))
}

func (sc *UserController) UpdateSelf(ctx *gin.Context) {
	id := utils.GetUserIdentity(ctx)
	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserInvalidIdentity, false, nil, nil))
		return
	}

	var commandRequest user_commands.UpdateUserCommandRequest
	if err := commandRequest.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserInvalidFields, false, []string{err.Error()}, nil))
		return
	}

	updaterID, _ := utils.GetUserID(ctx)
	commandRequest.Updater = updaterID
	res, err := sc.handler.UpdateSelf(id, commandRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.UserReadSuccess, true, nil, res))
}

func (sc *UserController) Delete(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		id := ctx.Param("id")
		pid, err := uuid.Parse(id)
		if err != nil {
			ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
			return
		}

		sid, err := utils.GetUserStudioID(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
			return
		}

		deleter, _ := utils.GetUserID(ctx)
		err = sc.handler.Delete(pid, sid, deleter)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserNotFound, false, []string{err.Error()}, nil))
			return
		}

		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.UserDeleteSuccess, true, nil, nil))
		return

	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
}

func (sc *UserController) Update(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		id := ctx.Param("id")
		pid, err := uuid.Parse(id)
		if err != nil {
			ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
			return
		}

		sid, err := utils.GetUserStudioID(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
			return
		}

		var command user_commands.UpdateUserCommandRequest
		if err := command.Bind(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserInvalidFields, false, []string{err.Error()}, nil))
			return
		}

		updaterID, _ := utils.GetUserID(ctx)
		command.Updater = updaterID

		res, err := sc.handler.Update(pid, sid, command)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserNotFound, false, []string{err.Error()}, nil))
			return
		}

		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.UserReadSuccess, true, nil, res))
		return
	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
}

func (sc *UserController) Read(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		id := ctx.Param("id")
		pid, err := uuid.Parse(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
			return
		}

		sid, err := utils.GetUserStudioID(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
			return
		}

		res, err := sc.handler.Read(pid, sid)
		if err != nil {
			ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.UserNotFound, false, []string{err.Error()}, nil))
			return
		}

		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.UserReadSuccess, true, nil, res))
		return
	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
}

func (sc *UserController) Upload(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		id := ctx.Param("id")
		pid, err := uuid.Parse(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
			return
		}

		sid, err := utils.GetUserStudioID(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
			return
		}

		file, handler, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.FileRetreiveFailed, true, nil, nil))
			return
		}
		defer file.Close()

		res, err := sc.handler.Upload(pid, sid, file, handler)
		if err != nil {
			ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.FileUploadFailed, false, []string{err.Error()}, nil))
			return
		}

		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.FileUploadSuccess, true, nil, res))
		return
	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
}

func (sc *UserController) UploadSelf(ctx *gin.Context) {
	id := utils.GetUserIdentity(ctx)
	if id == "" {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.UserInvalidIdentity, false, nil, nil))
		return
	}

	file, handler, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.FileRetreiveFailed, true, nil, nil))
		return
	}
	defer file.Close()

	res, err := sc.handler.UploadSelf(id, file, handler)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.FileUploadFailed, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.FileUploadSuccess, true, nil, res))
}
