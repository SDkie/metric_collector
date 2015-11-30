package db

import (
	"os"

	"github.com/SDkie/metric_collector/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db gorm.DB

func InitPg() {
	var err error
	db, err = gorm.Open("postgres", os.Getenv("PG_URL"))
	logger.PanicfIfError(err, "Error while connecting to PostgreSQL, %s", err)

	db.LogMode(true)
	db.SingularTable(true)
}

func GetPg() *gorm.DB {
	return &db
}

func ClosePg() {
	db.Close()
}
