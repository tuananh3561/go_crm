package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tuananh3561/go_crm/app/dto"
	"github.com/tuananh3561/go_crm/app/entity"
	"github.com/tuananh3561/go_crm/app/helper"
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
	role usecase.Role
}

func NewRoleController(
	role usecase.Role,
) RoleController {
	return &roleController{
		role: role,
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

	listAudioBook, err := r.role.RolesByParams(roleSearchDTO)
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

	gradeCreate, err := r.role.CreateRole(userParse, roleCreateDTO)
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

	err := r.role.UpdateRole(userParse, roleUpdateDTO)
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
