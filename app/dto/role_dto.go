package dto

type RoleSearchDTO struct {
	RoleId   string   `json:"role_id" form:"role_id"`
	RoleIds  []string `json:"role_ids" form:"role_ids"`
	Name     string   `json:"name" form:"name"`
	ParentId string   `json:"parent_id" form:"parent_id"`
	Status   int      `json:"status" form:"status"`
	Limit    int      `json:"limit" form:"limit"`
	Page     int      `json:"page" form:"page"`
}

type RoleCreateDTO struct {
	Id       string `json:"id" form:"id" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	ParentId string `json:"parent_id" form:"parent_id" default:""`
}

type RoleUpdateDTO struct {
	Id       string `json:"id" form:"id" validate:"required"`
	Name     string `json:"name" form:"name"`
	ParentId string `json:"parent_id" form:"parent_id" default:""`
}

type RoleUpdateStatusDTO struct {
	Id     string `json:"id" form:"id" validate:"required"`
	Status int    `json:"status" form:"status" validate:"required"`
}
