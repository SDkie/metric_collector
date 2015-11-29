package worker

import (
	"github.com/SDkie/metric_collector/logger"
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

func InitCron() {
	jobrunner.Start()

	err := jobrunner.Schedule("0 0 * * * *", WorkerHourlyLog{})
	logger.PanicfIfError(err, "Error while scheduling Worker for HourlyLog, %s", err)
}

func JobHtml(c *gin.Context) {
	// Returns the template data pre-parsed
	c.HTML(200, "", jobrunner.StatusPage())
}
