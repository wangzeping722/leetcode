package main

func uniqueMorseRepresentations(words []string) int {
	sb := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	mbytes := make(map[byte]string, 'z'-'a')
	for i := 0; i <= 'z'-'a'; i++ {
		mbytes[byte(i)+'a']=sb[i]
	}
	res := make(map[string]bool)
	for _, word := range words {
		mores := ""
		for _, b := range word {
			mores += mbytes[byte(b)]
		}
		res[mores] = true
	}

	return len(res)
}

func main() {
	uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})
}