package model

import (
	"strconv"

	"github.com/SDkie/metric_collector/db"
	"github.com/SDkie/metric_collector/logger"
	"github.com/garyburd/redigo/redis"
)

type MetricRedis struct {
	MetricStruct
}

func InitRedis() {
	db.InitRedis()
}

func (m *MetricRedis) Insert() error {
	// distinct_name:YYYY:MM:DD
	dailyBucket := "distinct_name:" + strconv.Itoa(m.CreatedAt.Year()) + ":" + strconv.Itoa(int(m.CreatedAt.Month())) + ":" + strconv.Itoa(m.CreatedAt.Day())
	_, err := db.GetRedisConnection().Do("SADD", dailyBucket, m.Metric)
	return err
}

func MergeToMonthlyBucket(dailyBucketKey, monthlyBucketKey string) {
	count, err := redis.Int(db.GetRedisConnection().Do("SCARD", dailyBucketKey))
	if count == 0 || err != nil {
		return
	}

	logger.Debugf("Merge Daily Bucket %s -> Monthly Bucket %s", dailyBucketKey, monthlyBucketKey)
	_, err = db.GetRedisConnection().Do("SUNIONSTORE", monthlyBucketKey, monthlyBucketKey, dailyBucketKey)
	if err != nil {
		logger.Errf("Error during SUNIONSTORE %s", err.Error())
		return
	}

	db.GetRedisConnection().Do("DEL", dailyBucketKey)
}
