// This is Go Main Func
package main // これは必ず記載する

import "fmt"

func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("bazz")
}

func main() {
	bazz()
	fmt.Println("Hello World") // 標準出力には fmt のimport が必要
}
