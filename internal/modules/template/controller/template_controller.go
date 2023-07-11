package template_controller

import (
	"net/http"

	"github.com/LucasToledoPereira/go-boilerplate/internal/datatypes"
	"github.com/LucasToledoPereira/go-boilerplate/internal/enums/codes"
	"github.com/LucasToledoPereira/go-boilerplate/internal/models"
	template_commands "github.com/LucasToledoPereira/go-boilerplate/internal/modules/template/commands"
	handler "github.com/LucasToledoPereira/go-boilerplate/internal/modules/template/handler"
	"github.com/LucasToledoPereira/go-boilerplate/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TemplateController struct {
	handler handler.TemplateHandler
}

func NewTemplateController(handler handler.TemplateHandler) (c TemplateController) {
	return TemplateController{
		handler: handler,
	}
}

// @Summary Create a new template for a game
// @Description Create a new template for a game using the provided request body
// @Tags Templates
// @Accept json
// @Produce json
// @Param idGame path string true "The ID of the game to create the template for"
// @Param request body CreateTemplateCommandRequest true "The request body containing the details of the template to be created"
// @Success 200 {object} NewResultWrapper[CreateTemplateCommandResponse]
// @Failure 400 {object} Error
// @Failure 401 {object} Error
// @Failure 500 {object} Error
// @Router /api/v1/templates/create/{idGame} [post]
func (sc *TemplateController) Create(ctx *gin.Context) {
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

	var commandRequest template_commands.CreateTemplateCommandRequest
	if err := commandRequest.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.TemplateInvalidFields, false, []string{err.Error()}, nil))
		return
	}

	creatorID, _ := utils.GetUserID(ctx)
	commandRequest.Creator = creatorID
	commandRequest.StudioID = studioId
	commandRequest.GameID = gameID

	msgs := commandRequest.Validate()
	if msgs != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.TemplateInvalidFields, false, msgs, nil))
		return
	}

	res, err := sc.handler.Create(commandRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.TemplateCreateFailed, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.TemplateCreateSuccess, true, nil, res))
}

func (sc *TemplateController) List(ctx *gin.Context) {
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
		ctx.JSON(http.StatusInternalServerError, models.NewResultWrapper[any](codes.TemplateListsFailed, false, []string{err.Error()}, nil))
		return
	}
	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.TemplateListsSuccess, true, nil, res))
}

func (sc *TemplateController) Read(ctx *gin.Context) {
	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	tid := ctx.Param("idTemplate")
	templateID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	res, err := sc.handler.Read(templateID, gameID, studioId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.TemplateNotFound, false, []string{err.Error()}, nil))
		return
	}
	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.TemplateReadSuccess, true, nil, res))
}

func (sc *TemplateController) Update(ctx *gin.Context) {
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

	tid := ctx.Param("idTemplate")
	templateID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioID, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	var command template_commands.UpdateTemplateCommandRequest
	if err := command.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.TemplateInvalidFields, false, []string{err.Error()}, nil))
		return
	}

	updaterID, _ := utils.GetUserID(ctx)
	command.Updater = updaterID
	command.StudioID = studioID
	command.GameID = gameID
	command.TemplateID = templateID

	res, err := sc.handler.Update(command)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.TemplateNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.TemplateUpdateSuccess, true, nil, res))
}

func (sc *TemplateController) Delete(ctx *gin.Context) {
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

	tid := ctx.Param("idTemplate")
	templateID, err := uuid.Parse(tid)
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
	err = sc.handler.Delete(templateID, gameID, studioId, deleter)

	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.TemplateNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.TemplateDeleteSuccess, true, nil, nil))
}

func (sc *TemplateController) CreateAttribute(ctx *gin.Context) {
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

	tid := ctx.Param("idTemplate")
	templateID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	var commandRequest template_commands.CreateTemplateAttributeCommandRequest
	if err := commandRequest.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.TemplateAttributeInvalidFields, false, []string{err.Error()}, nil))
		return
	}

	creatorID, _ := utils.GetUserID(ctx)
	commandRequest.Creator = creatorID
	commandRequest.StudioID = studioId
	commandRequest.GameID = gameID
	commandRequest.TemplateID = templateID

	msgs := commandRequest.Validate()
	if msgs != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.TemplateAttributeInvalidFields, false, msgs, nil))
		return
	}

	res, err := sc.handler.CreateAttribute(commandRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.TemplateAttributeCreateFailed, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.TemplateAttributeCreateSuccess, true, nil, res))
}

func (sc *TemplateController) UpdateAttribute(ctx *gin.Context) {
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

	tid := ctx.Param("idTemplate")
	templateID, err := uuid.Parse(tid)
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

	var command template_commands.UpdateTemplateAttributeCommandRequest
	if err := command.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.TemplateAttributeInvalidFields, false, []string{err.Error()}, nil))
		return
	}

	updaterID, _ := utils.GetUserID(ctx)
	command.Updater = updaterID
	command.AttributeID = attributeID
	command.TemplateID = templateID
	command.StudioID = studioId
	command.GameID = gameID

	res, err := sc.handler.UpdateAttribute(command)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.TemplateAttributeNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.TemplateAttributeUpdateSuccess, true, nil, res))
}

func (sc *TemplateController) ListAttributes(ctx *gin.Context) {
	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	tid := ctx.Param("idTemplate")
	templateID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	res, err := sc.handler.ListAttributes(gameID, studioId, templateID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewResultWrapper[any](codes.TemplateAttributeListsFailed, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.TemplateAttributeListsSuccess, true, nil, res))
}

func (sc *TemplateController) ReadAttribute(ctx *gin.Context) {
	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	tid := ctx.Param("idTemplate")
	templateID, err := uuid.Parse(tid)
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

	res, err := sc.handler.ReadAttribute(gameID, studioId, templateID, attributeID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.NewResultWrapper[any](codes.TemplateAttributeNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper(codes.TemplateAttributeReadSuccess, true, nil, res))
}

func (sc *TemplateController) DeleteAttribute(ctx *gin.Context) {
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

	tid := ctx.Param("idTemplate")
	templateID, err := uuid.Parse(tid)
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
	err = sc.handler.DeleteAttribute(attributeID, templateID, gameID, studioId, deleter)

	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.TemplateAttributeNotFound, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.TemplateAttributeDeleteSuccess, true, nil, nil))
}

func (sc *TemplateController) UploadImage(ctx *gin.Context) {
	sc.upload(ctx, datatypes.COVER)
}

func (sc *TemplateController) UploadFiles(ctx *gin.Context) {
	sc.upload(ctx, datatypes.ANY)
}

func (sc *TemplateController) ListFiles(ctx *gin.Context) {
	gid := ctx.Param("idGame")
	gameID, err := uuid.Parse(gid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	tid := ctx.Param("idTemplate")
	templateID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	res, err := sc.handler.ListFilesIgnoringCover(gameID, studioId, templateID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.FileListFailed, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.FileListSuccess, true, nil, res))
}

func (sc *TemplateController) upload(ctx *gin.Context, filetype datatypes.FileType) {
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

	tid := ctx.Param("idTemplate")
	templateID, err := uuid.Parse(tid)
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

	var command template_commands.UploadFileTemplateCommandRequest
	creatorID, _ := utils.GetUserID(ctx)
	command.Creator = creatorID
	command.File = file
	command.FileHeader = handler
	command.StudioID = studioId
	command.TemplateID = templateID
	command.GameID = gameID
	command.Type = filetype

	res, err := sc.handler.Upload(command)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.FileUploadFailed, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.FileUploadSuccess, true, nil, res))
}

func (sc *TemplateController) DeleteFiles(ctx *gin.Context) {
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

	tid := ctx.Param("idTemplate")
	templateID, err := uuid.Parse(tid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	fid := ctx.Param("idFile")
	fileID, err := uuid.Parse(fid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, true, nil, nil))
		return
	}

	studioId, err := utils.GetUserStudioID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResultWrapper[any](codes.InvalidParamID, false, nil, nil))
		return
	}

	err = sc.handler.DeleteFile(fileID, templateID, gameID, studioId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.NewResultWrapper[any](codes.FileDeleteFailed, false, []string{err.Error()}, nil))
		return
	}

	ctx.JSON(http.StatusOK, models.NewResultWrapper[any](codes.FileDeleteSuccess, true, nil, nil))
}
