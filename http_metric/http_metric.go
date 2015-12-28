package http_metric

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/SDkie/metric_collector/db"
	"github.com/SDkie/metric_collector/logger"
	"github.com/gin-gonic/gin"
)

//go:generate easytags http_metric.go bson

type HttpMetric struct {
	Id                string        `bson:"_id"`
	Count             int           `bson:"count"`
	TotalResponseTime time.Duration `bson:"total_response_time"`
}

var metric HttpMetric

const HTTP_METRIC = "HTTP_METRIC"

func Init() {
	hmetric := new(HttpMetric)
	err := db.MgFindOne(HTTP_METRIC, &bson.M{"_id": HTTP_METRIC}, hmetric)
	if err != nil {
		hmetric.Id = HTTP_METRIC
		hmetric.Count = 0
		hmetric.TotalResponseTime = 0
		db.MgCreate(HTTP_METRIC, hmetric)
	}
}

func HttpMetricMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		processingTime := time.Now().Sub(startTime)

		hmetric := new(HttpMetric)
		change := mgo.Change{
			Update:    bson.M{"$inc": bson.M{"count": 1, "total_response_time": processingTime}},
			ReturnNew: true,
		}

		/* MgFindAndModify does findAndModify atomically. So even if multiple
		instances of this is executed its data will always be consistent */
		err := db.MgFindAndModify(HTTP_METRIC, &bson.M{"_id": HTTP_METRIC}, change, hmetric)
		if err != nil {
			logger.Errf("Error while running MgFindAndModify %s", err)
			return
		}

		avg := time.Duration(int(hmetric.TotalResponseTime) / hmetric.Count)
		logger.Debugf("HTTP Metrics:\n No of Requests Handled: %d\n Current processing Time: %s\n Avg Processing Time: %s", hmetric.Count, processingTime, avg)
	}
}
