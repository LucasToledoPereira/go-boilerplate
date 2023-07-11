package collection_controller

import (
	"net/http"

	"github.com/LucasToledoPereira/go-boilerplate/internal/enums/codes"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	collection_commands "github.com/LucasToledoPereira/go-boilerplate/internal/modules/collection/commands"
	handler "github.com/LucasToledoPereira/go-boilerplate/internal/modules/collection/handler"
	"github.com/LucasToledoPereira/go-boilerplate/internal/utils"
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
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	var commandRequest collection_commands.CreateCollectionCommandRequest
	if err := commandRequest.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionInvalidFields, false, []string{err.Error()}, nil))
		return
	}

	creatorID, _ := utils.GetUserID(ctx)
	commandRequest.Creator = creatorID
	commandRequest.StudioID = studioId
	commandRequest.GameID = gameID

	msgs := commandRequest.Validate()
	if msgs != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionInvalidFields, false, msgs, nil))
		return
	}

	res, err := sc.handler.Create(commandRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionCreateFailed, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.CollectionCreateSuccess, true, nil, res))
}

func (sc *CollectionController) List(ctx *gin.Context) {
	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}
	res, err := sc.handler.List(gameID, studioId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewResultWrapper[any](codes.CollectionListsFailed, false, []string{err.Error()}, nil))
		return
	}
	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.CollectionListsSuccess, true, nil, res))
}

func (sc *CollectionController) Update(ctx *gin.Context) {
	//A Common User can not update templates
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioID, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	var command collection_commands.UpdateCollectionCommandRequest
	if err := command.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionInvalidFields, false, []string{err.Error()}, nil))
		return
	}

	updaterID, _ := utils.GetUserID(ctx)
	command.Updater = updaterID
	command.StudioID = studioID
	command.GameID = gameID
	command.CollectionID = collectionID

	res, err := sc.handler.Update(command)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.CollectionUpdateSuccess, true, nil, res))
}

func (sc *CollectionController) Read(ctx *gin.Context) {
	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	res, err := sc.handler.Read(collectionID, gameID, studioId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.CollectionNotFound, false, []string{err.Error()}, nil))
		return
	}
	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.CollectionReadSuccess, true, nil, res))
}

func (sc *CollectionController) Delete(ctx *gin.Context) {
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	deleter, _ := utils.GetUserID(ctx)
	err = sc.handler.Delete(collectionID, gameID, studioId, deleter)

	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.CollectionNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.TemplateDeleteSuccess, true, nil, nil))
}

func (sc *CollectionController) UploadAvatar(ctx *gin.Context) {
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	file, handler, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.FileRetreiveFailed, true, nil, nil))
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
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.FileUploadFailed, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.FileUploadSuccess, true, nil, res))
}

func (sc *CollectionController) UploadCover(ctx *gin.Context) {
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	file, handler, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.FileRetreiveFailed, true, nil, nil))
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
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.FileUploadFailed, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.FileUploadSuccess, true, nil, res))
}

func (sc *CollectionController) CreateAttribute(ctx *gin.Context) {
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	var commandRequest collection_commands.CreateCollectionAttributeCommandRequest
	if err := commandRequest.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionAttributeInvalidFields, false, []string{err.Error()}, nil))
		return
	}

	creatorID, _ := utils.GetUserID(ctx)
	commandRequest.Creator = creatorID
	commandRequest.StudioID = studioId
	commandRequest.GameID = gameID
	commandRequest.CollectionID = collectionID

	msgs := commandRequest.Validate()
	if msgs != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionAttributeInvalidFields, false, msgs, nil))
		return
	}

	res, err := sc.handler.CreateAttribute(commandRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionAttributeCreateFailed, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.CollectionAttributeCreateSuccess, true, nil, res))
}

func (sc *CollectionController) UpdateAttribute(ctx *gin.Context) {
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	aid := ctx.Param("idAttribute")
	attributeID, err := uuid.Parse(aid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	var command collection_commands.UpdateCollectionAttributeCommandRequest
	if err := command.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.CollectionAttributeInvalidFields, false, []string{err.Error()}, nil))
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
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.TemplateAttributeNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.TemplateAttributeUpdateSuccess, true, nil, res))
}

func (sc *CollectionController) ListAttributes(ctx *gin.Context) {
	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	res, err := sc.handler.ListAttributes(gameID, studioId, collectionID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewResultWrapper[any](codes.CollectionAttributeListsFailed, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.CollectionAttributeListsSuccess, true, nil, res))
}

func (sc *CollectionController) ReadAttribute(ctx *gin.Context) {
	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	aid := ctx.Param("idAttribute")
	attributeID, err := uuid.Parse(aid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	res, err := sc.handler.ReadAttribute(gameID, studioId, collectionID, attributeID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewResultWrapper[any](codes.CollectionAttributeNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.CollectionAttributeReadSuccess, true, nil, res))
}

func (sc *CollectionController) DeleteAttribute(ctx *gin.Context) {
	if utils.IsCommon(ctx) {
		ctx.JSON(http.StatusUnauthorized, models.NewResultWrapper[any](codes.NotAuthorized, false, nil, nil))
		return
	}

	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	tid := ctx.Param("idCollection")
	collectionID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	aid := ctx.Param("idAttribute")
	attributeID, err := uuid.Parse(aid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	deleter, _ := utils.GetUserID(ctx)
	err = sc.handler.DeleteAttribute(attributeID, collectionID, gameID, studioId, deleter)

	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.CollectionAttributeNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.CollectionAttributeDeleteSuccess, true, nil, nil))
}
