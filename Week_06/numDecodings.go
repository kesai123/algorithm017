package main

import "strconv"


func numDecodings(s string) int {
	a := []rune(s)
	dp := make([]int, len(a))

	for i := range a {
		if i==0 {
			dp[i] = count1String(a[i])
		} else if i==1 {
			dp[i] = dp[i-1]*count1String(a[i]) + count2String(a[i-1:i+1])
		} else {
			dp[i] = dp[i-1]*count1String(a[i]) + dp[i-2]*count2String(a[i-1:i+1])
		}
	}
	return dp[len(dp)-1]
}

func count1String(s rune) int{
	i, _ := strconv.Atoi(string(s))
	if i>=1 && i<=9 {return 1} else {return 0}
}

func count2String(s []rune) int{
	i, _ := strconv.Atoi(string(s))
	if i>=10 && i<=26 {return 1} else {return 0}
}
