package main

import "fmt"

func main(){
	s := "anagram"
	t := "nagaram"
	b1 := isAnagram(s,t)
	fmt.Printf("b1=%v\n", b1)
	s = "rat"
	t = "car"
	b2 := isAnagram(s,t)
	fmt.Printf("b2=%v\n", b2)

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var h = make(map[rune]int)
	for _, r := range s {
		c, exist := h[r]
		if exist {
			h[r] = c+1
		} else {
			h[r] = 1
		}
	}
	for _, r := range t {
		c, exist := h[r]
		if exist {
			h[r] = c-1
		} else {
			return false
		}
	}
	for _, v := range h {
		if v != 0 {
			return false
		}
	}
	return true
}