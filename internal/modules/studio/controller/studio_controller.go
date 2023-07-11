package studio_controller

import (
	"net/http"

	codes "github.com/LucasToledoPereira/go-boilerplate/internal/enums/codes"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	studio_commands "github.com/LucasToledoPereira/go-boilerplate/internal/modules/studio/commands"
	handler "github.com/LucasToledoPereira/go-boilerplate/internal/modules/studio/handler"
	"github.com/LucasToledoPereira/go-boilerplate/internal/utils"
	"github.com/gin-gonic/gin"
)

type StudioController struct {
	handler handler.StudioHandler
}

func NewStudioController(handler handler.StudioHandler) (c StudioController) {
	return StudioController{
		handler: handler,
	}
}

func (sc StudioController) Read(ctx *gin.Context) {
	id, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, []string{err.Error()}, nil))
		return
	}

	studio, err := sc.handler.Read(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.StudioNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.StudioReadSuccess, true, nil, studio))
}

func (sc StudioController) Delete(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		id, err := utils.GetUserStudioID(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, []string{err.Error()}, nil))
			return
		}
		err = sc.handler.Delete(id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.StudioDeleteFailed, false, []string{err.Error()}, nil))
			return
		}

		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.StudioDeleteSuccess, true, nil, nil))
		return
	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
}

func (sc StudioController) Update(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		id, err := utils.GetUserStudioID(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, []string{err.Error()}, nil))
			return
		}

		var commandRequest studio_commands.UpdateStudioCommandRequest
		if err = commandRequest.Bind(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.StudioInvalidFields, false, []string{err.Error()}, nil))
			return
		}

		res, err := sc.handler.Update(id, commandRequest)

		if err != nil {
			ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.StudioUpdateFailed, false, []string{err.Error()}, nil))
			return
		}

		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.StudioUpdateSuccess, true, nil, res))
		return
	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
}

func (sc StudioController) Upload(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		id, err := utils.GetUserStudioID(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, []string{err.Error()}, nil))
			return
		}

		file, handler, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.FileRetreiveFailed, false, []string{err.Error()}, nil))
			return
		}
		defer file.Close()

		res, err := sc.handler.Upload(id, file, handler)
		if err != nil {
			ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.FileUploadFailed, false, []string{err.Error()}, nil))
			return
		}

		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.FileUploadSuccess, true, nil, res))
		return
	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
}
