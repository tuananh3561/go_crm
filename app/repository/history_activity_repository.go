package repository

import (
	"context"
	"github.com/tuananh3561/go_crm/app/database/db/mongodb"
	"github.com/tuananh3561/go_crm/app/entity"
	"github.com/tuananh3561/go_crm/app/helper"
	"go.mongodb.org/mongo-driver/mongo"
)

type HistoryActivityRepository interface {
	CreateHistoryActivity(context context.Context, historyActivity entity.HistoryActivity) (entity.HistoryActivity, error)
}

type historyActivityConnection struct {
	client     *mongo.Client
	connection *mongo.Collection
}

func NewHistoryActivityRepository(dbConn *mongo.Client) HistoryActivityRepository {
	connection := mongodb.OpenCollection(dbConn, "history_activity")
	return &historyActivityConnection{
		client:     dbConn,
		connection: connection,
	}
}

//
//func (db HistoryActivity) GetListHistoryActivityByParams(ctx *gin.Context, params entity.HistoryActivitySearch) ([]entity.HistoryActivity, error) {
//	var historyActivities []entity.HistoryActivity
//	findOptions := options.Find()
//	if params.Limit != 0 {
//		offset := (params.Page - 1) * params.Limit
//		findOptions.SetLimit(int64(params.Limit))
//		findOptions.SetSkip(int64(offset))
//	}
//	filter := map[string]interface{}{}
//	filter = queryListHistoryActivityByParams(filter, params)
//	cursor, err := db.connection.Find(ctx, filter, findOptions)
//	for cursor.Next(ctx) {
//		var historyActivity model.HistoryActivity
//		errDecode := cursor.Decode(&historyActivity)
//		if errDecode != nil {
//			return nil, errDecode
//		}
//		historyActivities = append(historyActivities, model.MapDataHistoryActivity(historyActivity))
//	}
//	return historyActivities, err
//}
//
//func (db HistoryActivity) CountListHistoryActivityByParams(ctx *gin.Context, params entity.HistoryActivitySearch) (int64, error) {
//	filter := map[string]interface{}{}
//	filter = queryListHistoryActivityByParams(filter, params)
//	total, err := db.connection.CountDocuments(ctx, filter)
//	return total, err
//}
//
//func queryListHistoryActivityByParams(filter map[string]interface{}, params entity.HistoryActivitySearch) bson.M {
//	//if params.ActionController != "" {
//	//	filter["action"] = bson.M{"$regex": params.ActionController, "$options": "im"}
//	//}
//	if len(params.ListAction) > 0 {
//		filter["action"] = bson.M{"$in": params.ListAction}
//	}
//	if params.Action != "" {
//		filter["action"] = params.Action
//	}
//	if params.Id != 0 {
//		filter["condition.id"] = params.Id
//	}
//	if len(params.ListId) > 0 {
//		filter["condition.id"] = params.ListId
//	}
//	if params.IdUser != 0 {
//		filter["condition.idUser"] = params.IdUser
//	}
//	return filter
//}

func (db historyActivityConnection) CreateHistoryActivity(context context.Context, historyActivity entity.HistoryActivity) (entity.HistoryActivity, error) {
	historyActivity.CreatedAt = helper.TimeNow()
	_, err := db.connection.InsertOne(context, historyActivity)
	return historyActivity, err
}
