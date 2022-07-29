package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tuananh3561/go_crm/app/dto"
	"github.com/tuananh3561/go_crm/app/entity"
	"github.com/tuananh3561/go_crm/app/helper"
	"github.com/tuananh3561/go_crm/app/repository"
	"github.com/tuananh3561/go_crm/app/service"
	"github.com/tuananh3561/go_crm/app/usecase"
	"net/http"
)

type RoleController interface {
	List(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	UpdateStatus(context *gin.Context)
}

type roleController struct {
	roleRepo               repository.RoleRepository
	historyActivityService service.HistoryActivityService
}

func NewRoleController(
	roleRepo repository.RoleRepository,
	historyActivityService service.HistoryActivityService,
) RoleController {
	return &roleController{
		roleRepo:               roleRepo,
		historyActivityService: historyActivityService,
	}
}

func (r roleController) List(context *gin.Context) {
	message := "Get list roles "
	roleSearchDTO := dto.RoleSearchDTO{}

	errDTO := context.ShouldBind(&roleSearchDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse(message+"failed.", errDTO.Error(), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	listAudioBook, err := usecase.RolesByParams(r.roleRepo, roleSearchDTO)
	if err != nil {
		res := helper.BuildErrorResponse(message+"failed.", err.Error(), nil)
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildSuccessResponse(message+"success.", listAudioBook)
	context.JSON(http.StatusOK, res)
}

func (r roleController) Create(context *gin.Context) {
	message := "Create role "
	user, errGetUser := context.Get("user")
	if !errGetUser {
		res := helper.BuildErrorResponse(message+"failed.", "Not data user", nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	userParse := user.(entity.User)

	roleCreateDTO := dto.RoleCreateDTO{}

	errDTO := context.ShouldBind(&roleCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse(message+"failed.", errDTO.Error(), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if ok, err := helper.ValidateBase(roleCreateDTO); !ok {
		res := helper.BuildErrorResponse(message+"failed.", err, nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	gradeCreate, err := usecase.CreateRole(r.roleRepo, r.historyActivityService, userParse, roleCreateDTO)
	if err != nil {
		res := helper.BuildErrorResponse(message+"failed.", err.Error(), nil)
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildSuccessResponse(message+"success.", gradeCreate)
	context.JSON(http.StatusOK, res)
}

func (r roleController) Update(context *gin.Context) {
	message := "Update grade "
	user, errGetUser := context.Get("user")
	if !errGetUser {
		res := helper.BuildErrorResponse(message+"failed.", "Not data user", nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	userParse := user.(entity.User)

	roleUpdateDTO := dto.RoleUpdateDTO{}
	errDTO := context.ShouldBind(&roleUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse(message+"failed.", errDTO.Error(), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if ok, err := helper.ValidateBase(roleUpdateDTO); !ok {
		res := helper.BuildErrorResponse(message+"failed.", err, nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := usecase.UpdateRole(r.roleRepo, r.historyActivityService, userParse, roleUpdateDTO)
	if err != nil {
		res := helper.BuildErrorResponse(message+"failed.", err.Error(), nil)
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildSuccessResponse(message+"success.", nil)
	context.JSON(http.StatusOK, res)
}

func (r roleController) UpdateStatus(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}
