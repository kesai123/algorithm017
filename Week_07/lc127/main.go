package main

import "fmt"

func main(){
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot","dot","dog","lot","log","cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))

	beginWord = "hit"
	endWord = "cog"
	wordList = []string{"hot","dot","dog","lot","log"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
