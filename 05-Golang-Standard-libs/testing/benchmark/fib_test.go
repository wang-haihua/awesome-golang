package benchmark

import "testing"

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib2(1, b)
}

func BenchmarkFib2(b *testing.B) {
	benchmarkFib2(2, b)
}

func BenchmarkFib3(b *testing.B) {
	benchmarkFib2(3, b)
}
func BenchmarkFib10(b *testing.B) {
	benchmarkFib2(10, b)
}
func BenchmarkFib20(b *testing.B) {
	benchmarkFib2(20, b)
}
func BenchmarkFib40(b *testing.B) {
	benchmarkFib2(40, b)
}

func benchmarkFib2(n int, b *testing.B) {
	b.ReportAllocs() // ReportAllocs 方法用于打开当前基准测试的内存统计功能
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

/*
基准测试结果描述：
- 2000000 ：基准测试的迭代总次数 b.N
- 898 ns/op：平均每次迭代所消耗的纳秒数
- 368 B/op：平均每次迭代内存所分配的字节数
- 9 allocs/op：平均每次迭代的内存分配次数
*/
