package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.JSONFormatter{}

	now := time.Now().Format("2006.01.02.15:04:005")
	fileName := fmt.Sprintf("gin_%s.log", now)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("create log file error:%v", err)
		panic(err)
	}

	log.Out = file

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Out
	log.Level = logrus.InfoLevel
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args)
}
