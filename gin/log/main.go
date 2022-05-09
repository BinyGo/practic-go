package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	http := Router()
	//写入日志文件
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	// 如果你需要同时写入日志文件和控制台上显示，使用下面代码
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//gin.DisableConsoleColor() // 禁用日志的颜色
	//gin.ForceConsoleColor() //记录日志的颜色

	//
	http.Use(LoggerToFile())
	http.Use(gin.Recovery())
	http.Run(":8999")
}

func Router() *gin.Engine {

	router := gin.New() //默认的没有中间件的空白 Gin
	//router := gin.Default() // 默认已经连接了 Logger and Recovery 中间件
	router.GET("/ping", LoggerToFile(), func(ctx *gin.Context) {
		//Info级别的日志
		LoggerInfo().WithFields(logrus.Fields{
			"name": "biny",
		}).Info("记录一下日志", "Info")
		//Error级别的日志
		LoggerInfo().WithFields(logrus.Fields{
			"name": "biny",
		}).Error("记录一下日志", "Error")
		//Warn级别的日志
		LoggerInfo().WithFields(logrus.Fields{
			"name": "biny",
		}).Warn("记录一下日志", "Warn")
		//Debug级别的日志
		LoggerInfo().WithFields(logrus.Fields{
			"name": "biny",
		}).Debug("记录一下日志", "Debug")
		ctx.JSON(http.StatusOK, gin.H{"message": "biny"})
	})

	return router
}

func LoggerInfo() *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/info/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.TraceLevel)

	//设置日志格式
	// logger.SetFormatter(&logrus.TextFormatter{
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })

	// 换成json格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}

func Logger() *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.TraceLevel)

	//设置日志格式
	// logger.SetFormatter(&logrus.TextFormatter{
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })

	// 换成json格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}

func LoggerToFile() gin.HandlerFunc {
	logger := Logger()
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		// logger.Infof("| %3d | %13v | %15s | %s | %s |",
		// 	statusCode,
		// 	latencyTime,
		// 	clientIP,
		// 	reqMethod,
		// 	reqUri,
		// )

		// 日志变成json
		logger.WithFields(logrus.Fields{
			"status_code":      statusCode,
			"latency_time":     latencyTime,                                         //秒的换算:ms(毫秒),μs(微秒),ns(纳秒),ps(皮秒)
			"latency_time_str": strings.TrimSpace(fmt.Sprintf("%13v", latencyTime)), //秒的换算:ms(毫秒),μs(微秒),ns(纳秒),ps(皮秒)
			"client_ip":        clientIP,
			"req_method":       reqMethod,
			"req_uri":          reqUri,
		}).Info()
	}
}
