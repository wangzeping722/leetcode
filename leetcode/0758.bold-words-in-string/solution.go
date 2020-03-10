package main

import (
	"bytes"
	"strings"
)

func boldWords(words []string, S string) string {
	bold := make([]bool, len(S))
	for _, word := range words {
		n := strings.Index(S, word)
		for n != -1 {
			for i := n; i < n+len(word); i++ {
				bold[i] = true
			}
			t := strings.Index(S[n+1:], word)
			if t == -1 {
				break
			}
			n = t + n + 1
		}
	}

	buf := bytes.Buffer{}
	if bold[0] {
		buf.WriteString("<b>")
	}
	for i := 0; i < len(bold); i++ {
		buf.WriteByte(S[i])
		if i == len(bold)-1 {
			if bold[i] {
				buf.WriteString("</b>")
			}
			break
		}

		if !bold[i] && bold[i+1] {
			buf.WriteString("<b>")
		}

		if bold[i] && !bold[i+1] {
			buf.WriteString("</b>")
		}
	}
	println(buf.String())
	return buf.String()
}

func main() {
	boldWords([]string{"ccb", "b", "d", "cba", "dc"},
		"eeaadadadc")
}
