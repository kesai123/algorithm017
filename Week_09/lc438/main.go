package main

import "fmt"

func main() {
	s1 :="cbaebabacd"
	p1 := "abc"
	ret1 := findAnagrams(s1, p1)
	fmt.Println("ret1:", ret1)

	s2 :="abab"
	p2 := "ab"
	ret2 := findAnagrams(s2, p2)
	fmt.Println("ret2:", ret2)
}
