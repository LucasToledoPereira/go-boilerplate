package collection_controller

import (
	"net/http"

	"github.com/LucasToledoPereira/go-boilerplate/apis/game/enums/codes"
	collection_commands "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/collection/commands"
	handler "github.com/LucasToledoPereira/go-boilerplate/apis/game/modules/collection/handler"
	"github.com/LucasToledoPereira/go-boilerplate/apis/game/utils"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CollectionController struct {
	handler handler.CollectionHandler
}

func NewCollectionController(handler handler.CollectionHandler) (c CollectionController) {
	return CollectionController{
		handler: handler,
	}
}

func (sc *CollectionController) Create(ctx *gin.Context) {
	//A Common User can not create templates
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized.String(), false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, nil, nil))
		return
	}

	var commandRequest collection_commands.CreateCollectionCommandRequest
	if err := commandRequest.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionInvalidFields.String(), false, []string{err.Error()}, nil))
		return
	}

	creatorID, _ := utils.GetUserID(ctx)
	commandRequest.Creator = creatorID
	commandRequest.StudioID = studioId
	commandRequest.GameID = gameID

	msgs := commandRequest.Validate()
	if msgs != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionInvalidFields.String(), false, msgs, nil))
		return
	}

	res, err := sc.handler.Create(commandRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionCreateFailed.String(), false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.CollectionCreateSuccess.String(), true, nil, res))
}

func (sc *CollectionController) List(ctx *gin.Context) {
	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, nil, nil))
		return
	}
	res, err := sc.handler.List(gameID, studioId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewResultWrapper[any](codes.CollectionListsFailed.String(), false, []string{err.Error()}, nil))
		return
	}
	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.CollectionListsSuccess.String(), true, nil, res))
}

func (sc *CollectionController) Update(ctx *gin.Context) {
	//A Common User can not update templates
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized.String(), false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	studioID, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	var command collection_commands.UpdateCollectionCommandRequest
	if err := command.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionInvalidFields.String(), false, []string{err.Error()}, nil))
		return
	}

	updaterID, _ := utils.GetUserID(ctx)
	command.Updater = updaterID
	command.StudioID = studioID
	command.GameID = gameID
	command.CollectionID = collectionID

	res, err := sc.handler.Update(command)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionNotFound.String(), false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.CollectionUpdateSuccess.String(), true, nil, res))
}

func (sc *CollectionController) Read(ctx *gin.Context) {
	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, nil, nil))
		return
	}

	res, err := sc.handler.Read(collectionID, gameID, studioId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.CollectionNotFound.String(), false, []string{err.Error()}, nil))
		return
	}
	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.CollectionReadSuccess.String(), true, nil, res))
}

func (sc *CollectionController) Delete(ctx *gin.Context) {
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized.String(), false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, nil, nil))
		return
	}

	deleter, _ := utils.GetUserID(ctx)
	err = sc.handler.Delete(collectionID, gameID, studioId, deleter)

	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.CollectionNotFound.String(), false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.TemplateDeleteSuccess.String(), true, nil, nil))
}

func (sc *CollectionController) UploadAvatar(ctx *gin.Context) {
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized.String(), false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, nil, nil))
		return
	}

	file, handler, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.FileRetreiveFailed.String(), true, nil, nil))
		return
	}
	defer file.Close()

	var command collection_commands.UploadFileCollectionCommandRequest
	command.File = file
	command.FileHeader = handler
	command.StudioID = studioId
	command.CollectionID = collectionID
	command.GameID = gameID

	res, err := sc.handler.UploadAvatar(command)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.FileUploadFailed.String(), false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.FileUploadSuccess.String(), true, nil, res))
}

func (sc *CollectionController) UploadCover(ctx *gin.Context) {
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized.String(), false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, nil, nil))
		return
	}

	file, handler, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.FileRetreiveFailed.String(), true, nil, nil))
		return
	}
	defer file.Close()

	var command collection_commands.UploadFileCollectionCommandRequest
	command.File = file
	command.FileHeader = handler
	command.StudioID = studioId
	command.CollectionID = collectionID
	command.GameID = gameID

	res, err := sc.handler.UploadCover(command)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.FileUploadFailed.String(), false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.FileUploadSuccess.String(), true, nil, res))
}

func (sc *CollectionController) CreateAttribute(ctx *gin.Context) {
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized.String(), false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, nil, nil))
		return
	}

	var commandRequest collection_commands.CreateCollectionAttributeCommandRequest
	if err := commandRequest.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionAttributeInvalidFields.String(), false, []string{err.Error()}, nil))
		return
	}

	creatorID, _ := utils.GetUserID(ctx)
	commandRequest.Creator = creatorID
	commandRequest.StudioID = studioId
	commandRequest.GameID = gameID
	commandRequest.CollectionID = collectionID

	msgs := commandRequest.Validate()
	if msgs != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionAttributeInvalidFields.String(), false, msgs, nil))
		return
	}

	res, err := sc.handler.CreateAttribute(commandRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionAttributeCreateFailed.String(), false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.CollectionAttributeCreateSuccess.String(), true, nil, res))
}

func (sc *CollectionController) UpdateAttribute(ctx *gin.Context) {
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized.String(), false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	aid := ctx.Param("idAttribute")
	attributeID, err := uuid.Parse(aid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, nil, nil))
		return
	}

	var command collection_commands.UpdateCollectionAttributeCommandRequest
	if err := command.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionAttributeInvalidFields.String(), false, []string{err.Error()}, nil))
		return
	}

	updaterID, _ := utils.GetUserID(ctx)
	command.Updater = updaterID
	command.AttributeID = attributeID
	command.CollectionID = collectionID
	command.StudioID = studioId
	command.GameID = gameID

	res, err := sc.handler.UpdateAttribute(command)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.TemplateAttributeNotFound.String(), false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.TemplateAttributeUpdateSuccess.String(), true, nil, res))
}

func (sc *CollectionController) ListAttributes(ctx *gin.Context) {
	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, nil, nil))
		return
	}

	res, err := sc.handler.ListAttributes(gameID, studioId, collectionID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewResultWrapper[any](codes.CollectionAttributeListsFailed.String(), false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.CollectionAttributeListsSuccess.String(), true, nil, res))
}

func (sc *CollectionController) ReadAttribute(ctx *gin.Context) {
	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	aid := ctx.Param("idAttribute")
	attributeID, err := uuid.Parse(aid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, nil, nil))
		return
	}

	res, err := sc.handler.ReadAttribute(gameID, studioId, collectionID, attributeID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewResultWrapper[any](codes.CollectionAttributeNotFound.String(), false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.CollectionAttributeReadSuccess.String(), true, nil, res))
}

func (sc *CollectionController) DeleteAttribute(ctx *gin.Context) {
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized.String(), false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	aid := ctx.Param("idAttribute")
	attributeID, err := uuid.Parse(aid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID.String(), false, nil, nil))
		return
	}

	deleter, _ := utils.GetUserID(ctx)
	err = sc.handler.DeleteAttribute(attributeID, collectionID, gameID, studioId, deleter)

	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.CollectionAttributeNotFound.String(), false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.CollectionAttributeDeleteSuccess.String(), true, nil, nil))
}
