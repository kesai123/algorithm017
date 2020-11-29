package main

import "fmt"

func main(){
	inputs := []string{
		"42",
		"   -42",
		"4193 with words",
		"words and 987",
		"-91283472332",
		"   ",
		"+33",
	}
	for i:=range inputs{
		fmt.Println("input:", inputs[i], "output:", myAtoi(inputs[i]))
	}
}
