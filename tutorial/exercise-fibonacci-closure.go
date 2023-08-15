package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// どかどか足していくだけで良い
	// 関数がインスタンス化された時，内部に変数を持つので，それをもとにフィボナッチを計算すれば良い．
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
