package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
// Myreader っていう型が，Reader型になるためには，そいつがメソッドRead()を持たなきゃいけない．このメソッドは，
// Read(buffer []byte) (read_data_num int, err error)をつけていれば良い．だから，
// 与えられたバッファを全部Aで埋めて，その長さとnilを返せば良い．
func (r MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
