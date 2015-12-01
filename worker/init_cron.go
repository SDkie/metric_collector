package worker

import (
	"github.com/SDkie/metric_collector/logger"
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

func InitCron() {
	jobrunner.Start()

	err := jobrunner.Schedule("@every 1h", WorkerHourlyLog{})
	logger.PanicfIfError(err, "Error while scheduling Worker HourlyLog, %s", err)
	err = jobrunner.Schedule("@every 15m", WorkerAccountName{})
	logger.PanicfIfError(err, "Error while scheduling Worker AccountName, %s", err)
}

func JobHtml(c *gin.Context) {
	// Returns the template data pre-parsed
	c.HTML(200, "", jobrunner.StatusPage())
}
