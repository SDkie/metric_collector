package db

import (
	"os"

	"github.com/SDkie/metric_collector/logger"
	"github.com/garyburd/redigo/redis"
)

var conn redis.Conn

func InitRedis() {
	var err error
	conn, err = redis.Dial("tcp", os.Getenv("REDIS_URL"))
	logger.PanicfIfError(err, "Error while Initializing Redis, %s", err)
	logger.Info("Redis Successfully Initialize")
}

func GetRedisConnection() redis.Conn {
	return conn
}
