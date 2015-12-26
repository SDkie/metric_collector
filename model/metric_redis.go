package model

import (
	"time"

	"github.com/SDkie/metric_collector/db"
	"github.com/garyburd/redigo/redis"
)

func InitRedis() {
	db.InitRedis()

	// Initialize current_date Key in Redis
	_, err := redis.String(db.GetRedisConnection().Do("GET", "current_date"))
	if err != nil {
		db.GetRedisConnection().Do("SET", "current_date", time.Now().Format("2006:01:02"))
	}
}
