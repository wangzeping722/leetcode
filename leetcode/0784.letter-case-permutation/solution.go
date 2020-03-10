package _784_letter_case_permutation

var res []string
func letterCasePermutation(S string) []string {
	res = []string{}
	helper(S, "")
	return res
}

func helper(s string, str string) {
	if len(s) == 0 {
		res = append(res, str)
		return
	}
	helper(s[1:], str + string(s[0]))
	if s[0] >='a' && s[0] <= 'z' {
		helper(s[1:], str+string(s[0]+'A'-'a'))
	} else if s[0] >='A' && s[0] <= 'Z' {
		helper(s[1:], str+string(s[0]-'A'+'a'))
	}
}


