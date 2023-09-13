package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"jump/technical-case/internal/technical_case/api"
	"jump/technical-case/internal/util/connectors"
	"log"
	"os"
	"time"
)

func OpenLogFile(name string) (*os.File, error) {
	currentTime := time.Now()
	log_name := os.Getenv("PATH_LOG_FILE_TRANSACTION") + "app_" + name + "_" + currentTime.Format("2006_02_01") + ".log"
	f, err := os.OpenFile(log_name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0744)
	return f, err
}

func Gorm() {
	logfile, err := OpenLogFile("gorm")
	if err != nil {
		panic(err.Error())
	}
	loggerGorm := logger.New(log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Lshortfile), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info, // En PROD mettre le params à logger.Warn ou logger.Error
	})
	dsn := os.Getenv("DSN_MASTER")
	err = connectors.Set(postgres.Open(dsn), &gorm.Config{
		Logger: loggerGorm,
	})
	if err != nil {
		panic(err.Error())
	}
}

func Router() *gin.Engine {
	Gorm()
	router := gin.Default()
	logfile, err := OpenLogFile("gin")
	if err != nil {
		panic(err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(logfile)
	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("{\"time\":\"%s\",\"status_code\":\"%d\",\"latency\":\"%s\",\"client_ip\":\"%s\",\"method\":\"%s\",\"path\":\"%s\",\"protocole\":\"%s\",\"user_agent\":\"%s\",\"error_message\":\"%s\"}\n",
			param.TimeStamp.Format(time.RFC3339Nano),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.GET("/users", api.GetUsers)
	router.POST("/invoice", api.PostInvoice)
	router.POST("/transaction", api.PostTransaction)
	return router
}
