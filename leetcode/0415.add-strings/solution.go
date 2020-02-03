package _415_add_strings

func addStrings(num1 string, num2 string) string {
	maxLen := max(len(num1), len(num2))

	res := make([]byte, maxLen+1)
	var carry byte
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i,j,maxLen= i-1,j-1,maxLen-1{
		var n1, n2 byte
		if i >= 0 {
			n1 = num1[i]-'0'
		}
		if j >= 0 {
			n2 = num2[j]-'0'
		}

		tmp := n1 + n2 + carry
		carry = tmp/10
		res[maxLen] = tmp%10+'0'
	}

	if carry == 1 {
		res[0] = carry + '0'
		return string(res)
	}
	return string(res[1:])
}

func max(l1, l2 int) int {
	if l1 < l2 {
		return l2
	}
	return l1
}
