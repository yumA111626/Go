// This is Go Main Func
package main // これは必ず記載する

import (
	"fmt"
	"os/user"
	"time"
)

// import に関しての講義
func main() {
	fmt.Println("Hello World", time.Now()) // 標準出力には fmt のimport が必要
	fmt.Println(user.Current())
}
