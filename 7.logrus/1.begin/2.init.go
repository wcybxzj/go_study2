package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(log.WarnLevel)
}

//知识点1:设置log输出格式 json输出
//知识点2:控制输出级别
/*
{"level":"warning","msg":"The group's number increased tremendously!","number":122,"omg":true,"time":"2019-09-24T11:35:37+08:00"}
{"level":"fatal","msg":"The ice breaks!","number":100,"omg":true,"time":"2019-09-24T11:35:37+08:00"}
*/
func main() {
	//Info不会显示
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
