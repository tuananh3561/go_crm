package dto

type PermissionSearchDTO struct {
	RoleId   string   `json:"role_id" form:"role_id"`
	RoleIds  []string `json:"role_ids" form:"role_ids"`
	Name     string   `json:"name" form:"name"`
	ParentId string   `json:"parent_id" form:"parent_id"`
	Status   int      `json:"status" form:"status"`
	Limit    int      `json:"limit" form:"limit"`
	Page     int      `json:"page" form:"page"`
}

type PermissionCreateDTO struct {
	Id       string `json:"id" form:"id" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	Router   string `json:"router" form:"router"`
	Note     string `json:"note" form:"note"`
	Type     string `json:"type" form:"type" validate:"required"`
	IsPublic string `json:"is_public" form:"is_public" validate:"required"`
}

type PermissionUpdateDTO struct {
	Id       string `json:"id" form:"id" validate:"required"`
	Name     string `json:"name" form:"name"`
	Router   string `json:"router" form:"router"`
	Note     string `json:"note" form:"note"`
	Type     string `json:"type" form:"type"`
	IsPublic string `json:"is_public" form:"is_public"`
}

type PermissionUpdateStatusDTO struct {
	Id     string `json:"id" form:"id" validate:"required"`
	Status int    `json:"status" form:"status" validate:"required"`
}
