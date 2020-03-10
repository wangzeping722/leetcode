package main

import "strconv"

func ipToCIDR(ip string, n int) []string {
	res := []string{}
	sum := ipToInt(ip)
	for n > 0 {
		k := sum & -sum
		for k > n {
			k /= 2
		}
		n -= k
		res = append(res, intToIp(sum, k))
		sum += k
	}
	return res
}

func ipToInt(ip string) int {
	sum, num := 0, 0
	ip += "."
	for _, b := range ip {
		if b == '.' {
			sum = sum*256 + num
			num = 0
		} else {
			num = 10*num + int(b) - '0'
		}
	}
	return sum
}

func intToIp(sum, k int) string {
	ip := ""
	a := make([]int, 4)
	for i := 3; i >= 0; i-- {
		a[i] = sum % 256
		sum /= 256
	}

	for i := 0; i < 3; i++ {
		ip += strconv.Itoa(a[i]) + "."
	}
	tmp := -1
	for k > 0 {
		k /= 2
		tmp++
	}
	ip += strconv.Itoa(a[3]) + "/" + strconv.Itoa(32-tmp)
	return ip
}

func main() {
	ipToCIDR("255.0.0.7", 10)
}
