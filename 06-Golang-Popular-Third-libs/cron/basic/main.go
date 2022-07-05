package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	// 1. 创建cron对象，用这个对象管理定时任务
	c := cron.New()
	// 2. 调用cron对象的AddFunc()方法向管理器中添加定时任务
	c.AddFunc("@every 1s", clockPrint)
	c.AddFunc("@every 2s", func() {
		fmt.Println("tick per two seconds.")
	})

	// 3. 调用c.Start()启动定时循环
	c.Start()

	time.Sleep(time.Second * 5) // 防止主 goroutine 退出
}

func clockPrint() {
	fmt.Println("tick per second.")
}
