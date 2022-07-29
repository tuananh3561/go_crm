package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HistoryActivity struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	IdUser     int                `json:"user_id" bson:"user_id"`
	Email      string             `json:"email" bson:"email"`
	Action     string             `json:"action" bson:"action"`
	ActionNote string             `json:"action_note" bson:"action_note"`
	Condition  interface{}        `json:"condition" bson:"condition"`
	Data       interface{}        `json:"data" bson:"data"`
	CreatedAt  int                `json:"time_created" bson:"time_created"`
}

func (HistoryActivity) TableName() string {
	return "history_activity"
}
