package problem0020

func isValid1(s string) bool {
	if s == "" {
		return true
	}

	if len(s)%2 == 1 {
		return false
	}

	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		}

		if s[i] == ')' || s[i] == '}' || s[i] == ']' {
			if len(stack) > 0 {
				c := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if c == '(' && s[i] != ')' {
					return false
				}
				if c == '{' && s[i] != '}' {
					return false
				}
				if c == '[' && s[i] != ']' {

					return false
				}
			} else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}

	return true
}

func isValid(s string) bool {
	size := len(s)

	stack := make([]byte, size)

	top := 0
	for i := 0; i < size; i++ {
		c := s[i]
		switch c {
		case '(':
			stack[top] = c + 1
			top++
		case '[', '{':
			stack[top] = c + 2
			top++
		case ')', '}', ']':
			if top > 0 && stack[top-1] == c {
				top--
			} else {
				return false
			}

		}
	}
	return top == 0
}
