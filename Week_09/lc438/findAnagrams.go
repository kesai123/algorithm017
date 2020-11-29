package main

//滑动窗口
func findAnagrams(s string, p string) []int {
	ret := []int{}
	countP := make([]int, 26)
	countS := make([]int, 26)
	for _, v := range p {
		countP[v-'a']++
	}
	//左闭右开窗口
	left, right := 0, 0
	for right < len(s) {
		// 滑动右指针
		curR := s[right]-'a'
		countS[curR]++
		right++
		// 右指针计数超范围，滑动左指针
		for countS[curR] > countP[curR] {
			curL := s[left]-'a'
			countS[curL]--
			left++
		}
		if right-left == len(p) {
			ret = append(ret, left)
		}
	}
	return ret
}



//全排列解法，超时
/*
func findAnagrams(s string, p string) []int {
	var ret []int
	if len(s) < len(p) {return ret}
	// 获取p的组合集
	m := getPermutation(p)
	a := []rune(s)
	for i:=0; i<len(s)-len(p)+1; i++ {
		//判断s的子串sub是否在组合集中
		sub := string(a[i:i+len(p)])
		if _, ok := m[sub]; ok {
			ret = append(ret, i)
		}
	}
	return ret
}

//对p执行字符全排列，参照lc47
func getPermutation(p string) map[string]bool{
	m := map[string]bool{}
	a := []rune(p)
	dfs(a, m, 0)
	return m
}

func dfs(a []rune, result map[string]bool, level int) {
	if level == len(a)-1 {
		result[string(a)] = true
		return
	}
	unique := map[rune]bool {}
	for i:=level; i<len(a); i++ {
		if _, ok := unique[a[i]]; ok {continue}
		a[i], a[level] = a[level], a[i]
		dfs(a, result, level+1)
		a[i], a[level] = a[level], a[i]
		unique[a[i]] = true
	}
}*/