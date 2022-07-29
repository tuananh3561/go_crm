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

//type AudioBookSearchDTO struct {
//	IdApp       int    `json:"app_id" form:"app_id"`
//	IdAudioBook int    `json:"id_audio_book" form:"id_audio_book"`
//	Title       string `json:"title" form:"title"`
//	IdSeries    int    `json:"id_series" form:"id_series"`
//	IdGrade     int    `json:"id_grade" form:"id_grade"`
//	LevelSystem int    `json:"level_system" form:"level_system"`
//	IdUser      int    `json:"id_user" form:"id_user"`
//	Limit       int    `json:"limit" form:"limit"`
//	Page        int    `json:"page" form:"page"`
//}
//
//type AudioBookCreateDTO struct {
//	IdApp       int                   `json:"id_app" form:"id_app" validate:"required"`
//	IdLanguage  int                   `json:"id_language" form:"id_language" validate:"required"`
//	IdParent    int                   `json:"id_parent" form:"id_parent"`
//	Title       string                `json:"title" form:"title" validate:"required"`
//	Content     string                `json:"content" form:"content" validate:"required"`
//	Description string                `json:"description" form:"description"`
//	Extra       string                `json:"extra" form:"extra"`
//	IdSeries    int                   `json:"id_series" form:"id_series" validate:"required"`
//	IdGrade     int                   `json:"id_grade" form:"id_grade"  validate:"required"`
//	FileAudio   *multipart.FileHeader `json:"file_audio" form:"file_audio" validate:"required"`
//	FileThumb   *multipart.FileHeader `json:"file_thumb" form:"file_thumb" validate:"required"`
//}
//
//type AudioBookUpdateDTO struct {
//	IdAudioBook int                   `json:"id_audio_book" form:"id_audio_book" validate:"required"`
//	Title       string                `json:"title" form:"title"`
//	Content     string                `json:"content" form:"content"`
//	Description string                `json:"description" form:"description"`
//	Extra       string                `json:"extra" form:"extra"`
//	IdSeries    int                   `json:"id_series" form:"id_series"`
//	IdGrade     int                   `json:"id_grade" form:"id_grade"`
//	FileAudio   *multipart.FileHeader `json:"file_audio" form:"file_audio"`
//	FileThumb   *multipart.FileHeader `json:"file_thumb" form:"file_thumb"`
//}
//
//type AudioBookUpdateLevelSystemDTO struct {
//	IdAudioBook int    `json:"id_audio_book" form:"id_audio_book" validate:"required"`
//	TypeUser    int    `json:"type_user" form:"type_user"`
//	IdUser      int    `json:"id_user_assign" form:"id_user_assign"`
//	LevelSystem int    `json:"level_system" form:"level_system" validate:"required"`
//	Note        string `json:"note" form:"note"`
//}
//
//type CountAudioBookDTO struct {
//	IdAudioBook int `json:"id_audio_book" form:"id_audio_book"`
//}
//
//type ChangeStatusAudioBookDTO struct {
//	IdAudioBook int `json:"id_audio_book" form:"id_audio_book"`
//	Status      int `json:"status" form:"status"`
//}
