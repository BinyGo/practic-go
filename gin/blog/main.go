package main

import (
	"log"
	"net/http"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/practic-go/gin/blog/global"
	"github.com/practic-go/gin/blog/internal/model"
	"github.com/practic-go/gin/blog/internal/routers"
	"github.com/practic-go/gin/blog/pkg/logger"
	"github.com/practic-go/gin/blog/pkg/setting"
	"github.com/practic-go/gin/blog/pkg/tracer"
)

func init() {
	err := setupSetting() // 初始化配置
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupLogger() // 初始化日志
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupDBEngine() // 初始化数据库
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
}

func main() {
	router := routers.NewRouter()
	srv := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	srv.ListenAndServe()

}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   20, //单位 mb
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blog",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
