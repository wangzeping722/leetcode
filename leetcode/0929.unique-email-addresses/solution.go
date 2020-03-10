package main

import "strings"

func main() {
	numUniqueEmails([]string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"})
}

func numUniqueEmails(emails []string) int {
	mEmails := make(map[string]int)
	for _, email := range emails {
		strs := strings.Split(email, "@")
		strs[0] = strings.ReplaceAll(strs[0], ".", "")
		for i:=0; i<len(strs[0]); i++ {
			if strs[0][i] == '+' {
				strs[0] = strs[0][:i]
				break
			}
		}
		mEmails[strings.Join(strs, "@")]++
	}
	return len(mEmails)
}
