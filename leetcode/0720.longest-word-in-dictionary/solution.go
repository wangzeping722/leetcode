package main


import "sort"

func main() {
	longestWord1([]string{"w","wo","wor","worl","world"})
}

func longestWord(words []string) string {
	sort.Strings(words)
	m := make(map[string]bool, len(words))

	res := words[0]

	for _, w := range words {
		n := len(w)
		if n == 1 {
			m[w] = true
		} else if m[w[:n-1]] {
			m[w] = true
			if len(res) < len(w) {
				res = w
			}
		}
	}

	return res
}

func longestWord1(words []string) string {
	if len(words) == 0 {
		return ""
	}

	trie := &Trie{
		isString: false,
		Next:     [26]*Trie{},
	}

	for _, word := range words {
		trie.insert(word)
	}

	res := ""
	for _, word := range words {
		if trie.search(word) {
			if len(word) > len(res) {
				res = word
			} else if len(word) == len(res) && word < res {
				res = word
			}
		}
	}
	return res
}


type Trie struct {
	isString bool
	Next [26]*Trie
}

func (t *Trie) insert(word string) {
	root := t
	for _, r := range word {
		if root.Next[r-'a'] == nil {
			root.Next[r-'a'] = &Trie{
				isString:false,
				Next: [26]*Trie{},
			}
		}
		root = root.Next[r-'a']
	}
	root.isString = true
}


func (t *Trie) search(word string) bool {
	root := t
	for _, r := range word {
		if root.Next[r-'a'] == nil || root.Next[r-'a'].isString == false {
			return false
		}
		root = root.Next[r-'a']
	}
	return true
}

