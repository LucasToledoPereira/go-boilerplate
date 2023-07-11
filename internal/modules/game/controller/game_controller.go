package game_controller

import (
	"net/http"

	"github.com/LucasToledoPereira/go-boilerplate/internal/enums/codes"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	game_commands "github.com/LucasToledoPereira/go-boilerplate/internal/modules/game/commands"
	handler "github.com/LucasToledoPereira/go-boilerplate/internal/modules/game/handler"
	"github.com/LucasToledoPereira/go-boilerplate/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GameController struct {
	handler handler.GameHandler
}

func NewGameController(handler handler.GameHandler) (c GameController) {
	return GameController{
		handler: handler,
	}
}

func (sc GameController) Create(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		studioId, err := utils.GetUserStudioID(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
			return
		}

		var commandRequest game_commands.CreateGameCommandRequest
		if err := commandRequest.Bind(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.GameInvalidFields, false, []string{err.Error()}, nil))
			return
		}

		creatorID, _ := utils.GetUserID(ctx)
		commandRequest.Creator = creatorID
		commandRequest.StudioID = studioId
		msgs := commandRequest.Validate()
		if msgs != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.GameInvalidFields, false, msgs, nil))
			return
		}

		res, err := sc.handler.Create(commandRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.GameCreateFailed, false, []string{err.Error()}, nil))
			return
		}

		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.GameCreateSuccess, true, nil, res))
		return
	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
}

func (sc GameController) List(ctx *gin.Context) {
	studioID, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewResultWrapper[any](codes.InvalidParamID, false, []string{err.Error()}, nil))
		return
	}
	res, err := sc.handler.List(studioID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewResultWrapper[any](codes.GameListsFailed, false, []string{err.Error()}, nil))
		return
	}
	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.GameListsSuccess, true, nil, res))
}

func (sc GameController) Read(ctx *gin.Context) {
	id := ctx.Param("idGame")
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
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.GameNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.GameReadSuccess, true, nil, res))
}

func (sc GameController) Update(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		id := ctx.Param("idGame")
		gameID, err := uuid.Parse(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
			return
		}

		studioID, err := utils.GetUserStudioID(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
			return
		}

		var command game_commands.UpdateGameCommandRequest
		if err := command.Bind(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.GameInvalidFields, false, []string{err.Error()}, nil))
			return
		}

		updaterID, _ := utils.GetUserID(ctx)
		command.Updater = updaterID
		command.StudioID = studioID
		command.GameID = gameID

		res, err := sc.handler.Update(command)
		if err != nil {
			ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.GameNotFound, false, []string{err.Error()}, nil))
			return
		}

		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.GameUpdateSuccess, true, nil, res))
		return
	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
}

func (sc GameController) Delete(ctx *gin.Context) {
	if utils.IsOwnerOrAdministrator(ctx) {
		id := ctx.Param("idGame")
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

		deleter, _ := utils.GetUserID(ctx)
		err = sc.handler.Delete(pid, sid, deleter)

		if err != nil {
			ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.GameNotFound, false, []string{err.Error()}, nil))
			return
		}

		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.GameDeleteSuccess, true, nil, nil))
		return

	}
	ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
}
