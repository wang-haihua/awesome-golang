package benchmark

func Fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 0
	}
	return Fib(n-1) + Fib(n-2)
}
