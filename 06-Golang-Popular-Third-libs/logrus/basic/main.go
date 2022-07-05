package main

import "github.com/sirupsen/logrus"

/*
logrus的使用非常简单，与标准库log类似。logrus支持更多的日志级别：

	Panic：记录日志，然后panic。
	Fatal：致命错误，出现错误时程序无法正常运转。输出日志后，程序退出；
	Error：错误日志，需要查看原因；
	Warn：警告信息，提醒程序员注意；
	Info：关键操作，核心流程的日志；
	Debug：一般程序中输出的调试信息；
	Trace：很细粒度的信息，一般用不到；

日志级别从上向下依次增加，Trace最大，Panic最小。logrus有一个日志级别，高于这个级别的日志不会输出。默认的级别为InfoLevel。所以为了能看到Trace和Debug日志，我们在main函数第一行设置日志级别为TraceLevel。
*/

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{}) // 设置日志格式

	logrus.Trace("trace message content")
	logrus.Error("error message content")
	logrus.Warn("warn message content")
	logrus.Info("info message content")
	logrus.Debug("debug message content")
	logrus.Fatal("fatal message content") // 致命错误会导致程序退出, 执行完成之后调用 os.exit(1)
	logrus.Panic("panic message content") // 执行完后调用 panic 方法抛出异常信息
}
