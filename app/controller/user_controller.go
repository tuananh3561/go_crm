package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	UsersByParams(context *gin.Context)
	CreateUser(context *gin.Context)
	UpdateUser(context *gin.Context)
}

type userController struct {
}

func (u userController) UsersByParams(context *gin.Context) {
	//message := "Get list audio books "
	//audioBookSearchDTO := dto.UserSearchDTO{}
	//
	//errDTO := context.ShouldBind(&audioBookSearchDTO)
	//if errDTO != nil {
	//	res := helper.BuildErrorResponse(message+"failed.", errDTO.Error(), nil)
	//	context.AbortWithStatusJSON(http.StatusBadRequest, res)
	//	return
	//}
	//
	//audioBookSearch := entity.AudioBookSearch{
	//	IdApp:       audioBookSearchDTO.IdApp,
	//	IdAudioBook: audioBookSearchDTO.IdAudioBook,
	//	Title:       audioBookSearchDTO.Title,
	//	IdGrade:     audioBookSearchDTO.IdGrade,
	//	IdSeries:    audioBookSearchDTO.IdSeries,
	//	LevelSystem: audioBookSearchDTO.LevelSystem,
	//	IdUser:      audioBookSearchDTO.IdUser,
	//	IsAssign:    true,
	//	Limit:       audioBookSearchDTO.Limit,
	//	Page:        audioBookSearchDTO.Page,
	//}
	//
	//listAudioBook, err := usecase.GetListAudioBookByParams(c.audioBookRepo, audioBookSearch)
	//if err != nil {
	//	res := helper.BuildErrorResponse(message+"failed.", err.Error(), nil)
	//	context.AbortWithStatusJSON(http.StatusInternalServerError, res)
	//	return
	//}
	//
	//res := helper.BuildSuccessResponse(message+"success.", listAudioBook)
	//context.JSON(http.StatusOK, res)
	//TODO implement me
	panic("implement me")
}

func (u userController) CreateUser(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u userController) UpdateUser(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewAudioBookController() UserController {
	return &userController{}
}

//func (c audioBookController) GetListAudioBookByParams(context *gin.Context) {
//
//}
//
//func (c audioBookController) CreateAudioBook(context *gin.Context) {
//	message := "Create audio book "
//	user, errGetUser := context.Get("user")
//	if !errGetUser {
//		res := helper.BuildErrorResponse(message+"failed.", "Not data user", nil)
//		context.AbortWithStatusJSON(http.StatusBadRequest, res)
//		return
//	}
//	userParse := user.(entity.User)
//
//	audioBookCreateDTO := dto.AudioBookCreateDTO{}
//	errDTO := context.ShouldBind(&audioBookCreateDTO)
//	if errDTO != nil {
//		res := helper.BuildErrorResponse(message+"failed.", errDTO.Error(), nil)
//		context.AbortWithStatusJSON(http.StatusBadRequest, res)
//		return
//	}
//
//	if ok, errValidate := helper.ValidateBase(audioBookCreateDTO); !ok {
//		res := helper.BuildErrorResponse(message+"failed.", errValidate, nil)
//		context.AbortWithStatusJSON(http.StatusBadRequest, res)
//		return
//	}
//
//	fileNameAudio, duration, errUploadMediaAudio := usecase.UploadFileAudio(context, c.mediaService, audioBookCreateDTO.FileAudio)
//	if errUploadMediaAudio != nil {
//		res := helper.BuildErrorResponse(message+"failed.", errUploadMediaAudio.Error(), nil)
//		context.AbortWithStatusJSON(http.StatusBadRequest, res)
//		return
//	}
//	fileNameThumb, errUploadMediaThumb := usecase.UploadFileThumb(context, c.mediaService, audioBookCreateDTO.FileThumb)
//	if errUploadMediaThumb != nil {
//		res := helper.BuildErrorResponse(message+"failed.", errUploadMediaThumb.Error(), nil)
//		context.AbortWithStatusJSON(http.StatusBadRequest, res)
//		return
//	}
//
//	fileSizeThumb := float32(audioBookCreateDTO.FileAudio.Size)
//	audioBook := entity.AudioBook{
//		IdApp:       audioBookCreateDTO.IdApp,
//		IdLanguage:  audioBookCreateDTO.IdLanguage,
//		IdParent:    audioBookCreateDTO.IdParent,
//		Title:       &audioBookCreateDTO.Title,
//		Description: &audioBookCreateDTO.Description,
//		Content:     &audioBookCreateDTO.Content,
//		Extra:       &audioBookCreateDTO.Extra,
//		IdSeries:    &audioBookCreateDTO.IdSeries,
//		IdGrade:     &audioBookCreateDTO.IdGrade,
//		Thumb:       &fileNameThumb,
//		Audio:       &fileNameAudio,
//		Duration:    &duration,
//		AudioSize:   &fileSizeThumb,
//	}
//
//	response, err := usecase.CreateAudioBook(c.audioBookRepo, c.seriesRepo, c.gradeRepo, audioBook, userParse)
//	if err != nil {
//		res := helper.BuildErrorResponse(message+"failed.", err.Error(), nil)
//		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
//		return
//	}
//
//	res := helper.BuildSuccessResponse(message+"success.", response)
//	context.JSON(http.StatusOK, res)
//}
//
//func (c audioBookController) UpdateAudioBook(context *gin.Context) {
//	message := "Update audio book "
//	user, errGetUser := context.Get("user")
//	if !errGetUser {
//		res := helper.BuildErrorResponse(message+"failed.", "Not data user", nil)
//		context.AbortWithStatusJSON(http.StatusBadRequest, res)
//		return
//	}
//	userParse := user.(entity.User)
//
//	audioBookUpdateDTO := dto.AudioBookUpdateDTO{}
//	errDTO := context.ShouldBind(&audioBookUpdateDTO)
//	if errDTO != nil {
//		res := helper.BuildErrorResponse(message+"failed.", errDTO.Error(), nil)
//		context.AbortWithStatusJSON(http.StatusBadRequest, res)
//		return
//	}
//
//	if ok, errValidate := helper.ValidateBase(audioBookUpdateDTO); !ok {
//		res := helper.BuildErrorResponse(message+"failed.", errValidate, nil)
//		context.AbortWithStatusJSON(http.StatusBadRequest, res)
//		return
//	}
//
//	audioBook := entity.AudioBook{
//		IdAudioBook: audioBookUpdateDTO.IdAudioBook,
//		Title:       &audioBookUpdateDTO.Title,
//		Content:     &audioBookUpdateDTO.Content,
//		Description: &audioBookUpdateDTO.Description,
//		Extra:       &audioBookUpdateDTO.Extra,
//		IdSeries:    &audioBookUpdateDTO.IdSeries,
//		IdGrade:     &audioBookUpdateDTO.IdGrade,
//	}
//
//	if audioBookUpdateDTO.FileAudio != nil {
//		fileNameAudio, duration, errUploadMediaAudio := usecase.UploadFileAudio(context, c.mediaService, audioBookUpdateDTO.FileAudio)
//		if errUploadMediaAudio != nil {
//			res := helper.BuildErrorResponse(message+"failed.", errUploadMediaAudio.Error(), nil)
//			context.AbortWithStatusJSON(http.StatusBadRequest, res)
//			return
//		}
//		audioBook.Audio = &fileNameAudio
//		audioBook.Duration = &duration
//	}
//
//	if audioBookUpdateDTO.FileThumb != nil {
//		fileNameThumb, errUploadMediaThumb := usecase.UploadFileThumb(context, c.mediaService, audioBookUpdateDTO.FileThumb)
//		if errUploadMediaThumb != nil {
//			res := helper.BuildErrorResponse(message+"failed.", errUploadMediaThumb.Error(), nil)
//			context.AbortWithStatusJSON(http.StatusBadRequest, res)
//			return
//		}
//		audioBook.Thumb = &fileNameThumb
//	}
//
//	err := usecase.UpdateAudioBook(c.audioBookRepo, c.seriesRepo, c.gradeRepo, c.assignRepo, audioBook, userParse)
//	if err != nil {
//		res := helper.BuildErrorResponse(message+"failed.", err.Error(), nil)
//		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
//		return
//	}
//
//	res := helper.BuildSuccessResponse(message+"success.", nil)
//	context.JSON(http.StatusOK, res)
//}
//
//func (c audioBookController) UpdateLevelSystemAudioBook(context *gin.Context) {
//	message := "Update level system audio book "
//	user, errGetUser := context.Get("user")
//	if !errGetUser {
//		res := helper.BuildErrorResponse(message+"failed.", "Not data user", nil)
//		context.AbortWithStatusJSON(http.StatusBadRequest, res)
//		return
//	}
//	userParse := user.(entity.User)
//
//	audioBookUpdateLevelSystemDTO := dto.AudioBookUpdateLevelSystemDTO{}
//	errDTO := context.ShouldBind(&audioBookUpdateLevelSystemDTO)
//	if errDTO != nil {
//		res := helper.BuildErrorResponse(message+"failed.", errDTO.Error(), nil)
//		context.AbortWithStatusJSON(http.StatusBadRequest, res)
//		return
//	}
//
//	if ok, errValidate := helper.ValidateBase(audioBookUpdateLevelSystemDTO); !ok {
//		res := helper.BuildErrorResponse(message+"failed.", errValidate, nil)
//		context.AbortWithStatusJSON(http.StatusBadRequest, res)
//		return
//	}
//	audioBookUpdateLevelSystem := entity.AudioBookUpdateLevelSystem{
//		IdAudioBook: audioBookUpdateLevelSystemDTO.IdAudioBook,
//		TypeUser:    audioBookUpdateLevelSystemDTO.TypeUser,
//		IdUser:      audioBookUpdateLevelSystemDTO.IdUser,
//		LevelSystem: audioBookUpdateLevelSystemDTO.LevelSystem,
//		Note:        audioBookUpdateLevelSystemDTO.Note,
//	}
//	err := usecase.UpdateLevelSystemAudioBook(c.audioBookRepo, c.assignRepo, audioBookUpdateLevelSystem, userParse)
//	if err != nil {
//		res := helper.BuildErrorResponse(message+"failed.", err.Error(), nil)
//		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
//		return
//	}
//
//	res := helper.BuildSuccessResponse(message+"success.", nil)
//	context.JSON(http.StatusOK, res)
//}
