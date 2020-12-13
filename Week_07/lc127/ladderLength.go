package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	sol := NewSolution(beginWord, endWord, wordList)
	if sol == nil {return 0}
	return sol.BidirectionalBFS()
}

type Solution struct{
	//单个串长度
	wordLen int
	//合法字符串字典
	wordSet map[string]bool
	//字符串变化去重
	visitedSet map[string]bool
	//起点队列，通过一次性替换模拟BTS队列更新
	beginSet map[string]bool
	//终点队列
	endSet map[string]bool
}

func NewSolution(beginWord string, endWord string, wordList []string) *Solution {
	sol := Solution{
		len(beginWord),
		map[string]bool{},
		map[string]bool{},
		map[string]bool{beginWord:true},
		map[string]bool{endWord:true},
	}
	//初始化合法字典
	for i := range wordList {
		sol.wordSet[wordList[i]] = true
	}
	//目标字符不在字典中，提前返回失败
	if _, ok := sol.wordSet[endWord]; !ok {return nil}
	return &sol
}

//执行双向BFS
func (sol *Solution)BidirectionalBFS() int {
	step := 1
	for len(sol.beginSet)>0 && len(sol.endSet)>0 {
		// 对较小队列执行BTS，始终将begin限定为小队列，简化处理流程
		if len(sol.beginSet)>len(sol.endSet){
			sol.beginSet, sol.endSet = sol.endSet, sol.beginSet
		}
		// 下一层BTS队列
		nextLayerSet := map[string]bool{}
		//处理当前层队列中每个string
		for word := range sol.beginSet {
			if sol.IsValidWord(word, nextLayerSet) {return step+1}
		}
		//更新执行队列，进入下一层BTS
		sol.beginSet = nextLayerSet
		step++
	}
	return 0
}

//判断当前词是满足接龙要求，更新BTS队列
func (sol *Solution)IsValidWord(word string, nextLayerSet map[string]bool) bool{
	chs := []rune(word)
	// 穷举当前单词进行变化，依次变换每个位置的每个字符
	for j:=0;j<sol.wordLen;j++ {
		org := chs[j]
		for c:='a';c<='z';c++{
			if org == c {continue}
			chs[j]=c
			//当前词在终点队列，则双向BTS相遇，接龙完成
			if _,ok := sol.endSet[string(chs)]; ok {return true}
			//跳过不在字典中的词
			if _,ok := sol.wordSet[string(chs)]; !ok {continue}
			// 将变化后的词放入下一层队列
			if _,ok := sol.visitedSet[string(chs)]; !ok {
				nextLayerSet[string(chs)] = true
				sol.visitedSet[string(chs)] = true
			}
		}
		chs[j] = org
	}
	return false
}


/*
func ladderLength1(beginWord string, endWord string, wordList []string) int {

	step := 1
	//可选路径set
	wordSet := map[string]bool{}
	for i := range wordList {
		wordSet[wordList[i]] = true
	}
	if _, ok := wordSet[endWord]; !ok {return 0}

	//已访问set
	visitedSet := map[string]bool{}
	//起始、终止set
	beginSet := map[string]bool{beginWord:true}
	endSet := map[string]bool{endWord:true}

	//双向BTS
	for len(beginSet)>0 && len(endSet)>0 {
		// 优先对较小Set执行BTS
		if len(beginSet)>len(endSet){
			beginSet, endSet = endSet, beginSet
		}
		// 下一层BTS set
		nextLayerSet := map[string]bool{}
		//处理当前层set中的每个string
		for s := range beginSet {
			chs := []rune(s)
			// 穷举当前单词进行变化
			for j:=0;j<len(beginWord);j++ {
				org := chs[j]
				for c:='a';c<='z';c++{
					if org == c {continue}
					chs[j]=c
					//变化的单次在endSet中，则双向BTS set相遇，返回结果
					if _,ok := endSet[string(chs)]; ok {return step+1}
					//跳过不合法字符串
					if _,ok := wordSet[string(chs)]; !ok {continue}
					// 将当前词的变化放入下一层Set
					if _,ok := visitedSet[string(chs)]; !ok {
						nextLayerSet[string(chs)] = true
						visitedSet[string(chs)] = true
					}
				}
				chs[j] = org
			}
		}
		//进入下一层BTS
		beginSet = nextLayerSet
		step++
	}
	return 0
}*/