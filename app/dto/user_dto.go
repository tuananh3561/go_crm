package dto

type UserSearchDTO struct {
	IdApp       int    `json:"app_id" form:"app_id"`
	IdAudioBook int    `json:"id_audio_book" form:"id_audio_book"`
	Title       string `json:"title" form:"title"`
	IdSeries    int    `json:"id_series" form:"id_series"`
	IdGrade     int    `json:"id_grade" form:"id_grade"`
	LevelSystem int    `json:"level_system" form:"level_system"`
	IdUser      int    `json:"id_user" form:"id_user"`
	Limit       int    `json:"limit" form:"limit"`
	Page        int    `json:"page" form:"page"`
}

type UserCreateDTO struct {
	Id       string `json:"id" form:"id" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	ParentId string `json:"parent_id" form:"parent_id" default:""`
}

type UserUpdateDTO struct {
	Id       string `json:"id" form:"id" validate:"required"`
	Name     string `json:"name" form:"name"`
	ParentId string `json:"parent_id" form:"parent_id" default:""`
}

type UserUpdateStatusDTO struct {
	Id     string `json:"id" form:"id" validate:"required"`
	Status int    `json:"status" form:"status" validate:"required"`
}
