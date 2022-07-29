package service

import "github.com/tuananh3561/go_crm/app/entity"

type HistoryActivityService interface {
	CreateHistoryActivity(user entity.User, action string, actionNote string, condition interface{}, data interface{})
}

type historyActivityService struct {
}

func NewHistoryActivityService() HistoryActivityService {
	return &historyActivityService{}
}

func (h historyActivityService) CreateHistoryActivity(user entity.User, action string, actionNote string, condition interface{}, data interface{}) {
	_ = entity.HistoryActivity{
		IdUser:     user.Id,
		Email:      user.Email,
		Action:     action,
		ActionNote: actionNote,
		Condition:  condition,
		Data:       data,
	}
	//_, _ = h.historyActivityRepo.CreateHistoryActivity(context.TODO(), historyActivity)
}
