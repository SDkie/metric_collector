package model

import "github.com/SDkie/metric_collector/db"

//go:generate easytags metric_pg.go json
//go:generate easytags metric_pg.go sql

type MetricPg struct {
	Id int64 `sql:"id" gorm:"primary_key" json:"id"`
	MetricStruct
}

func InitPg() {
	db.InitPg()
	db.GetPg().CreateTable(&MetricPg{})
}

func (m *MetricPg) Insert() error {
	return db.GetPg().Create(m).Error
}
