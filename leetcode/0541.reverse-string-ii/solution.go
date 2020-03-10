package _541_reverse_string_ii

func reverseStr(s string, k int) string {
	ls := len(s)
	bytes := []byte(s)
	n := ls / (2*k)
	for i := 0; i < n; i++ {
		for m, n := 2*k*i, (2*i+1)*k-1; m < n; m, n = m+1, n-1 {
			bytes[m], bytes[n] = bytes[n], bytes[m]
		}
	}

	n = ls % (2*k)
	left := ls-n
	right := 0
	if n < k {
		right = ls-1
	} else {
		right = left+k-1
	}
	for ; left < right; left, right = left+1, right-1 {
		bytes[left], bytes[right] = bytes[right], bytes[left]
	}
	return string(bytes)
}
