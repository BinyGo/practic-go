package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

// logrus提供了New()函数来创建一个logrus的实例.
// 项目中,可以创建任意数量的logrus实例.
var log2 = logrus.New()

func Demo2() {
	// 为当前logrus实例设置消息的输出,同样地,
	// 可以设置logrus实例的输出到任意io.writer
	log2.Out = os.Stdout
	// 为当前logrus实例设置消息输出格式为json格式.
	// 同样地,也可以单独为某个logrus实例设置日志级别和hook,这里不详细叙述.

	log2.Formatter = &logrus.JSONFormatter{}
	log2.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log2.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log2.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}
