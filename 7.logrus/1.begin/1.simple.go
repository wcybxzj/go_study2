package main

import (
	"github.com/sirupsen/logrus"
)

//输出普通文字:
//time="2019-09-24T11:26:03+08:00" level=info msg="A walrus appears" animal=walrus

/*
	logrus.SetReportCaller(true)后的输出

	time="2019-09-24T14:31:40+08:00"
	level=info msg="A walrus appears"
	func=main.test1 file="/Users/ybx/www/go_www2/zuji/go_study2/7.logrus/1.begin/1.simple.go:14"
	animal=walrus
 */
func test1()  {
	//增加了函数调用位置
	logrus.SetReportCaller(true)
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}

func test2()  {
	logrus.Debug("Useful debugging information.")
	logrus.Info("Something noteworthy happened!")
	logrus.Warn("You should probably take a look at this.")
	logrus.Error("Something failed but I'm not quitting.")
	logrus.Fatal("Bye.")   //log之后会调用os.Exit(1)
	logrus.Panic("I'm bailing.")   //log之后会panic()
}

func main() {
	test1()
	//test2()
}