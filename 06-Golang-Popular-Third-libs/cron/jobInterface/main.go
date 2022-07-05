package main

import (
	"06-Golang-Popular-Third-libs/cron/jobInterface/models"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New()

	// 调用cron对象的AddJob()方法将GreetingJob对象添加到定时管理器中
	c.AddJob("@every 1s", models.GreetingJob{"Job-1", "001"})

	c.Start()

	time.Sleep(time.Second * 5)
}
