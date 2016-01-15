package main

import("fmt")

func anagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) { return false }

	chars1 := make(map[string]int)
	chars2 := make(map[string]int)

	for i:=0; i<len(s1); i++ {
		chars1[string(s1[i])] = chars1[string(s1[i])] + 1
	}
	for i:=0; i<len(s2); i++ {
		chars2[string(s2[i])] = chars2[string(s2[i])] + 1
	}

	keys := make([]string, len(chars1))
	for k := range chars1 {
		keys = append(keys, k)
	}

	for i:=0; i<len(keys); i++ {
		key := keys[i]
		if chars1[string(key)] != chars2[string(key)] { return false }
	}

	return true
}

func main(){
	fmt.Println(anagram("alejo", "ojela"))
	fmt.Println(anagram("aaallleeejooo", "aaallleeejooo"))
}
