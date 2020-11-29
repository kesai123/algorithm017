package main

import "math"

func myAtoi(s string) int {
	// range for result
	i, sign, sum := 0, 1, 0
	a := []rune(s)
	//skip ' '
	for ;i < len(a); i++{
		if a[i] != ' ' {break}
	}
	if i == len(a){return sum}
	//get sign
	if a[i] == '-' {
		sign = -1
		i++
	} else if a[i] == '+' {
		i++
	}
	//calculate
	for ;i < len(a); i++{
		if a[i] < '0' || a[i] > '9' {break}
		sum = sum*10 + int(a[i]-'0')
		// check range
		if sum * sign >= math.MaxInt32 {
			return math.MaxInt32
		} else if sum * sign <= math.MinInt32 {
			return math.MinInt32
		}
	}
	return sum*sign
}