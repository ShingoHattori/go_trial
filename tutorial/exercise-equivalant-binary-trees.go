package main

import "golang.org/x/tour/tree"

import (
	"fmt"	
)


// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// とりあえずもらって，左辺と右辺があるならそいつらをprint，いいえ．与えられた木がnilかどうかを見た方が良い
	// なければ，自分が保持している値をprintでどう？
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	// 二つの木をとりあえずWalkして，chに流れてくるデータが違ったらNGってことでどうよ
	t1ch, t2ch := make(chan int), make(chan int)

	// このように，処理が終わったら適切にcloseしなきゃいけない．そうじゃないと，デッドロックが起こる．
	go func() {
		defer close(t1ch)
		Walk(t1, t1ch)
	}()
	go func() {
		defer close(t2ch)
		Walk(t2, t2ch)
	}()
	
	//go Walk(t1, t1ch)
	//go Walk(t2, t2ch)
	for {
		val1, ok1 := <- t1ch
		val2, ok2 := <- t2ch
		if (!ok1 || !ok2) {
			return ok1 == ok2	
		}
		if (val1 != val2) {
			return false
		}
		
	}
}

func main() {
	treesize:= 10
	bintree := tree.New(treesize)
	readch  := make(chan int)
	go Walk(bintree, readch)
	for i:= 0; i<treesize; i++ {
		fmt.Println(<-readch)	
	}
}
