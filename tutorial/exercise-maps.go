package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// mapってのはつまるところ辞書なのである．
func WordCount(s string) map[string]int {
	// これで，sをスペースで分割して，スライスにする
	words := strings.Fields(s)

	// 戻り値
	wordCount := make(map[string]int)
	
	for _, word := range(words) {
		// wordCountの方は，未知のキーが来たらnullが入るんじゃなくて，宣言した時に入ってた型のゼロ値，つまり今はintなので0が入る．
		wordCount[word] ++	
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
