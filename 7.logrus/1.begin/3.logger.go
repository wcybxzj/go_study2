package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

//知识点1:将logrus做成外部对象
//知识点2:设置日志输出位置

// logger是一种相对高级的用法, 对于一个大型项目, 往往需要一个全局的logrus实例，即logger对象来记录项目所有的日志。
// logrus提供了New()函数来创建一个logrus的实例。
// 项目中，可以创建任意数量的logrus实例。
var log3 = logrus.New()

//输出:
/*
{
	"animal":"walrus",
	"level":"info",
	"msg":"A group of walrus emerges from the ocean",
	"size":10,
	"time":"2019-09-24T11:45:39+08:00"
}
*/
func main() {
 	// 为当前logrus实例设置消息的输出，同样地，
	// 可以设置logrus实例的输出到任意io.writer
	log3.Out = os.Stdout

	// 为当前logrus实例设置消息输出格式为json格式。
	// 同样地，也可以单独为某个logrus实例设置日志级别和hook，这里不详细叙述。
	log3.Formatter = &logrus.JSONFormatter{}

	log3.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}