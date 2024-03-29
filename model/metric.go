package model

import "time"

//go:generate easytags metric.go json
//go:generate easytags metric.go bson
//go:generate easytags metric.go sql

type MetricStruct struct {
	Username string `sql:"username;unique" bson:"username" json:"username"`
	Count    int64  `sql:"count" bson:"count" json:"count"`
	Metric   string `sql:"metric" bson:"metric" json:"metric"`

	CreatedAt time.Time `sql:"created_at" bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `sql:"updated_at" bson:"updated_at" json:"updated_at"`
	DeletedAt time.Time `sql:"deleted_at" bson:"deleted_at" json:"-"`
}
