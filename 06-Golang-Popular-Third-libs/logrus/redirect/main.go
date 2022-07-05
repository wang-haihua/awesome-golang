package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

func main() {
	writer1 := &bytes.Buffer{}                                                            // 缓冲区
	writer2 := os.Stdout                                                                  // 输出到标准输出
	writer3, err := os.OpenFile("logrus/redirect/log.txt", os.O_WRONLY|os.O_CREATE, 0755) // 输出到指定文件
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3)) // 多个输出
	logrus.Info("my testing log of multiple writers")
}
