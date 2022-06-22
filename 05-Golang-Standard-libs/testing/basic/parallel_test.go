package basic

import "testing"

var pairs = []struct {
	key string
	val string
}{
	{"one", "1"},
	{"two", "2"},
	{"three", "3"},
	{"four", "4"},
	{"five", "5"},
	{"six", "6"},
	{"seven", "7"},
	{"eight", "8"},
	{"nine", "9"},
}

func TestWriteToMap(t *testing.T) {
	t.Parallel()
	for _, pair := range pairs {
		WriteToMap(pair.key, pair.val)
	}
}

func TestReadFromMap(t *testing.T) {
	t.Parallel()
	for _, pair := range pairs {
		actual := ReadFromMap(pair.key)
		if actual != pair.val {
			t.Errorf("got the value of key(%s) is %s; expected is %s", pair.key, actual, pair.val)
		}
	}
}

/*
试验步骤：
1-注释掉 WriteToMap 和 ReadFromMap 中 locker 保护的代码，同时注释掉测试代码中的 t.Parallel，执行测试，测试通过，即使加上 -race，测试依然通过；
2-只注释掉 WriteToMap 和 ReadFromMap 中 locker 保护的代码，执行测试，测试失败（如果未失败，加上 -race 一定会失败）；
如果代码能够进行并行测试，在写测试时，尽量加上 Parallel，这样可以测试出一些可能的问题。
*/
