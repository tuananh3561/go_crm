package service

import (
	"github.com/tuananh3561/go_crm/app/entity"
	"github.com/tuananh3561/go_crm/app/job/rabbit_mq"
	"github.com/tuananh3561/go_crm/app/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HistoryActivityService interface {
	CreateHistoryActivity(user entity.User, action string, actionNote string, condition interface{}, data interface{})
}

type historyActivityService struct {
	historyActivityRepo repository.HistoryActivityRepository
}

func NewHistoryActivityService(
	historyActivityRepo repository.HistoryActivityRepository,
) HistoryActivityService {
	return &historyActivityService{
		historyActivityRepo: historyActivityRepo,
	}
}

func (h historyActivityService) CreateHistoryActivity(user entity.User, action string, actionNote string, condition interface{}, data interface{}) {
	historyActivity := entity.HistoryActivity{
		ID:         primitive.NewObjectID(),
		IdUser:     user.Id,
		Email:      user.Email,
		Action:     action,
		ActionNote: actionNote,
		Condition:  condition,
		Data:       data,
	}
	rabbit_mq.Publish(historyActivity)
	//_, _ = h.historyActivityRepo.CreateHistoryActivity(context.TODO(), historyActivity)
}
