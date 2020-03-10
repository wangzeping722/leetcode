package _509_fib


func fib(N int) int {
	if N <= 1 {
		return N
	}

	f1, f2 := 0, 1

	for i := 2; i < N; i++ {
		f := f1 + f2
		f2, f1 = f, f2
	}

	return f1 + f2
}