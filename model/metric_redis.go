package model

import (
	"time"

	"strconv"

	"github.com/SDkie/metric_collector/db"
	"github.com/garyburd/redigo/redis"
)

type MetricRedis struct {
	Metric
}

func InitRedis() {
	db.InitRedis()

	// Initialize current_date Key in Redis
	_, err := redis.String(db.GetRedisConnection().Do("GET", "current_date"))
	if err != nil {
		db.GetRedisConnection().Do("SET", "current_date", time.Now().Format("2006:01:02"))
	}
}

func (m *MetricRedis) Insert() error {
	dailyBucket := "distinct_name:" + strconv.Itoa(m.CreatedAt.Day())
	_, err := db.GetRedisConnection().Do("SADD", dailyBucket, m.Metric)
	return err
}
