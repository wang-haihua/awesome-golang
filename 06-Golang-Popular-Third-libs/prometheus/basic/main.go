package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

/*
OPTS用于绑定被创建的大多数 Metric 类型的选项。

Name 字段必须设置为非空字符串表示 Metric 类型的名称。
所有其他字段都是可选的，可以安全地保留它们的零值。
通常 Help 字段也需要设置为非空字符串。

type Opts struct {
	Namespace string
	Subsystem string
	Name      string

	// Help provides information about this metric
	Help string

	ConstLabels Labels
}
*/

var (
	// Cpu 温度监控指标 Gauge 类型 测量值
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})

	// 硬盘失效次数监控指标 Counter 类型
	hardDiskFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hd_errors_total",
			Help: "Number of hard-disk errors.",
		},
		[]string{"device"}, // 附加了一个标签以将其转换为向量
	)
)

func init() {
	// 注册需要暴露的 Metrics
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hardDiskFailures)
}

func main() {
	cpuTemp.Set(65.3) // Set() 设置cupTemp为任意值
	hardDiskFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9091", nil)) // ListenAndServe监听srv.Addr指定的TCP地址，并且会调用Serve方法接收到的连接。如果srv.Addr为空字符串，会使用":http"。
}
