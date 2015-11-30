package model

import (
	"github.com/SDkie/metric_collector/db"
	"gopkg.in/mgo.v2/bson"
)

//go:generate easytags metric_mongo.go json
//go:generate easytags metric_mongo.go bson

const METRIC_MONGO_COLLECTION = "metric_collector"

type MetricMongo struct {
	Id     bson.ObjectId `bson:"_id" json:"id"`
	Metric `bson:",inline"`
}

func InitMongo() {
	db.InitMongo()
}

func (m *MetricMongo) Insert() error {
	return db.MgCreate(METRIC_MONGO_COLLECTION, *m)
}
