package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.WithFields(logrus.Fields{
		"name": "node-1",
		"id":   "001",
	}).Info("withField Info log")

	logrus.Info("withoutField Info log")

	// 使用WithFields的返回值定义一个带字段的日志输出方式
	myLogger := logrus.WithFields(logrus.Fields{
		"userId":    10001,
		"ip":        "192.168.73.22",
		"collector": "zeno",
	})

	myLogger.Trace("trace log content")
	myLogger.Debug("debug log content") // 只有 Info 和 Error 类型的日志输出了 这是因为什么？
	myLogger.Info("info log content")
	myLogger.Error("error log content")
}
