package _811_subdomain_visit_count

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	domainsCount := make(map[string]int)
	for _, domain := range cpdomains {
		ss := strings.Split(domain, " ")
		num, _ := strconv.Atoi(ss[0])
		subDomains := strings.Split(ss[1], ".")
		subDomain := ""
		for i := len(subDomains) - 1; i >= 0; i-- {
			if i != len(subDomains)-1 {
				subDomain = "." + subDomain
			}
			subDomain = subDomains[i] + subDomain
			domainsCount[subDomain] += num
		}
	}
	res := []string{}
	for k, v := range domainsCount {
		cpdomain := fmt.Sprintf("%d %s", v, k)
		res = append(res, cpdomain)
	}
	return res
}
