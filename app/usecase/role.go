package usecase

import (
	"errors"
	"github.com/tuananh3561/go_crm/app/dto"
	"github.com/tuananh3561/go_crm/app/entity"
	"github.com/tuananh3561/go_crm/app/repository"
	"github.com/tuananh3561/go_crm/app/service"
	"go.mongodb.org/mongo-driver/bson"
)

type Role interface {
	RolesByParams(params dto.RoleSearchDTO) (interface{}, error)
	CreateRole(user entity.User, roleCreateDTO dto.RoleCreateDTO) (*entity.Role, error)
	UpdateRole(user entity.User, roleUpdateDTO dto.RoleUpdateDTO) error
}

type user struct {
	roleRepo repository.RoleRepository
	history  service.HistoryActivityService
}

func NewRole(roleRepo repository.RoleRepository, history service.HistoryActivityService) *user {
	return &user{
		roleRepo: roleRepo,
		history:  history,
	}
}

func (u user) RolesByParams(params dto.RoleSearchDTO) (interface{}, error) {
	data, errData := u.roleRepo.RolesByParams(params)
	if params.Limit != 0 {
		total, errTotal := u.roleRepo.CountRoleByParams(params)
		response := dto.ResponseList{
			Data:  data,
			Total: total,
		}
		return response, errTotal
	}
	return data, errData
}

func (u user) CreateRole(user entity.User, roleCreateDTO dto.RoleCreateDTO) (*entity.Role, error) {
	roleCreate := entity.Role{
		Id:       roleCreateDTO.Id,
		Name:     roleCreateDTO.Name,
		ParentId: roleCreateDTO.ParentId,
		Status:   entity.RoleStatusActive,
	}
	role, errCreate := u.roleRepo.CreateRole(roleCreate)
	if errCreate == nil {
		//condition := bson.M{
		//	"id": roleCreateDTO.Id,
		//}
		//u.history.CreateHistoryActivity(user, entity.ListAction["grade"]["create"], "", condition, roleCreate)
	}
	return role, errCreate
}

func (u user) UpdateRole(user entity.User, roleUpdateDTO dto.RoleUpdateDTO) error {
	if roleUpdateDTO.Id == "" {
		return errors.New("id role not found")
	}
	role, err := u.roleRepo.FindRoleById(roleUpdateDTO.Id)
	if err != nil || role.Id == "" {
		return errors.New("id role " + roleUpdateDTO.Id + " does not exist")
	}
	roleUpdate := entity.Role{
		Id:       roleUpdateDTO.Id,
		Name:     roleUpdateDTO.Name,
		ParentId: roleUpdateDTO.ParentId,
	}
	errUpdate := u.roleRepo.UpdateRole(roleUpdate)
	if errUpdate == nil {
		condition := bson.M{
			"id": roleUpdate.Id,
		}
		u.history.CreateHistoryActivity(user, "role@update", "", condition, roleUpdate)
	}
	return errUpdate
}
