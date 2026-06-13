// This is Go Main Func
package main // これは必ず記載する

import "fmt" // 数値型 に関しての講義
func main() {
	// Go では長い物に合わせて記載する
	var (
		u8   uint8   = 221
		f32  float32 = 1.01
		int2 int     = 344
	)
	fmt.Printf("type=%T , Value=%v\n", u8, u8)
	fmt.Printf("type=%T , Value=%v\n", f32, f32)
	fmt.Printf("type=%T , Value=%v\n", int2, int2)
}
