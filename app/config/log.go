package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tuananh3561/go_crm/app/helper"
	"io"
	"os"
	"time"
)

const PathLog = "storage/logs"

func SetupLogOutput() {
	// create found log
	errMkd := helper.Mkdir(PathLog)
	if errMkd != nil {
		panic(errMkd)
	}
	// get file name log
	fileName := fmt.Sprintf(PathLog+"/gin_%s.log", time.Now().Format(time.ANSIC))
	// create file
	f, _ := os.Create(fileName)
	// save log
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] | %s   %s   %d   %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	})
}
