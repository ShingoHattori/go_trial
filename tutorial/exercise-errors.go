import (
	"fmt"
)

// ErrNegativeSqrtで，エラーになったfloatを返す
type ErrNegativeSqrt float64

// ErrNegativeSqrtに，Error()っていうメソッドをつける．こいつは文字列を返したりする．
func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf("cannnot Sqrt negative number : %v", float64(e))
	}
	return ""
}

// returnで返すものは，表示したいやつと，Error()メソッドを持った型．このメソッドを持ってる変数は，Printlnが自動でそれを見に行って，ErrorだったらError扱いしてくれる．	
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	z := 1.
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)	
	}
	return z, nil
}




func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	
}
