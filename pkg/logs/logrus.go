package logs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})

	// 设置将日志输出到指定文件（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	//log.SetOutput(os.Stdout) 终端输出
	logFile := "logrus.log"
	file, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(file) //写入到文件

	// 设置只记录日志级别为warn及其以上的日志
	log.SetLevel(log.WarnLevel)
}

func Demo1() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}

//更多logrus使用方式,logrus hook 存储等
//可参考:https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247484320&idx=1&sn=d6c4b46c7e106ca14368c0772b8c803a&chksm=fa80d237cdf75b2194539edf548abad7aaba64b998d3abab6d441b15fdbbaf1eddd258c0c30e&cur_album_id=1323498303014780929&scene=189#wechat_redirect
