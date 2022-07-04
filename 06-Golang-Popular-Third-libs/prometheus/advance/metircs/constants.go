package metircs

import (
	"06-Golang-Popular-Third-libs/prometheus/advance/random"
	"fmt"
	"net/http"
	"strconv"
)

const (
	namePrefix = "the_number_of_student"
	subSys     = "client_golang"
	nameSpace  = "prometheus_demo"
)

var names = []string{"小明", "小华", "小红"}
var ages = []int64{18, 20, 22, 23, 26, 28}

func GetParamNum(req *http.Request) int64 {
	err := req.ParseForm()
	if err != nil {
		fmt.Errorf("parse form error")
		return 0
	}
	numStr := req.Form.Get("num")
	fmt.Println("numStr:[%v]\n", numStr)
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		fmt.Errorf("parse int error :%v\n", err)
		return 0
	}
	return num
}

// randData随机获取一个元素，并且将本次请求的随机元素统计到countStrMap
func getCurRandomStrMap(countStrMap map[string]int64, randData []string) string {
	index := random.RandomNum(0, int64(len(randData)))
	randValue := randData[index]
	countStrMap[randValue] = countStrMap[randValue] + 1
	return randValue
}

// randData随机获取一个元素，并且将本次请求的随机元素统计到countIntMap
func getCurRandomIntMap(countIntMap map[int64]int64, randData []int64) string {
	index := random.RandomNum(0, int64(len(randData)))
	randVal := randData[index]
	countIntMap[randVal] = countIntMap[randVal] + 1
	return fmt.Sprintf("%d", randVal)
}
