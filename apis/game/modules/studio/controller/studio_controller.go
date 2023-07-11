package studio_controller

import (
	"net/http"

	codes "github.com/LucasToledoPereira/go-boilerplate/apis/game/enums/codes"
	studio_commands "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/studio/commands"
	handler "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/studio/handler"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/utils"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
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
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, []string{err.Error()}, nil))
		return
	}

	studio, err := sc.handler.Read(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.StudioNotFound.String(), false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.StudioReadSuccess.String(), true, nil, studio))
}

func (sc StudioController) Delete(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		id, err := utils.GetUserStudioID(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, []string{err.Error()}, nil))
			return
		}
		err = sc.handler.Delete(id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.StudioDeleteFailed.String(), false, []string{err.Error()}, nil))
			return
		}

		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.StudioDeleteSuccess.String(), true, nil, nil))
		return
	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized.String(), false, nil, nil))
}

func (sc StudioController) Update(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		id, err := utils.GetUserStudioID(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, []string{err.Error()}, nil))
			return
		}

		var commandRequest studio_commands.UpdateStudioCommandRequest
		if err = commandRequest.Bind(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.StudioInvalidFields.String(), false, []string{err.Error()}, nil))
			return
		}

		res, err := sc.handler.Update(id, commandRequest)

		if err != nil {
			ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.StudioUpdateFailed.String(), false, []string{err.Error()}, nil))
			return
		}

		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.StudioUpdateSuccess.String(), true, nil, res))
		return
	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized.String(), false, nil, nil))
}

func (sc StudioController) Upload(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		id, err := utils.GetUserStudioID(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, []string{err.Error()}, nil))
			return
		}

		file, handler, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.FileRetreiveFailed.String(), false, []string{err.Error()}, nil))
			return
		}
		defer file.Close()

		res, err := sc.handler.Upload(id, file, handler)
		if err != nil {
			ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.FileUploadFailed.String(), false, []string{err.Error()}, nil))
			return
		}

		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.FileUploadSuccess.String(), true, nil, res))
		return
	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized.String(), false, nil, nil))
}
