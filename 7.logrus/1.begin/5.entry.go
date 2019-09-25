package main

import log "github.com/sirupsen/logrus"

//输出
//INFO[0000] something happened on that request            request_id=111 user_ip=100.100.100.100
//WARN[0000] something not great happened                  request_id=111 user_ip=100.100.100.100


/*
前面的WithFields API可以规范使用者按照其提倡的方式记录日志。
但是WithFields依然是可选的，因为某些场景下，使用者确实只需要记录仪一条简单的消息。

通常，在一个应用中、或者应用的一部分中，都有一些固定的Field。
比如在处理用户http请求时，上下文中，所有的日志都会有request_id和user_ip。
为了避免每次记录日志都要使用log.WithFields(log.Fields{"request_id": request_id, "user_ip": user_ip})，
我们可以创建一个logrus.Entry实例，为这个实例设置默认Fields，在上下文中使用这个logrus.Entry实例记录日志即可。
*/
func main() {
	requestLogger := log.WithFields(log.Fields{"request_id": 111, "user_ip": "100.100.100.100"})
	requestLogger.Info("something happened on that request")
	requestLogger.Warn("something not great happened")
}
