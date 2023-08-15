package main

import (
	"io"
	"os"
	"strings"
)

// こういう型があって，こういう型の変数を必要としている
type rot13Reader struct {
	r io.Reader
}

// Reader型としてやっていけるように実装する
func (r rot13Reader) Read(b []byte) (n int, err error) {
	// とりあえず文字を読み出す
	n, err = r.r.Read(b) 
	// １文字ずつずらしていく
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])	
	}
	return n, err
}

func rot13(b byte) byte {
	switch {
	case 'a' <= b && b <= 'z':
		return 'a' + (b-'a'+13)%26
	case 'A' <= b && b <= 'Z':
		return 'A' + (b-'A'+13)%26
	default:
		return b
	}	
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
