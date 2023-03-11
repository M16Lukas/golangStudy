// パッケージは1つだけ宣言可能
package main

// パッケージをインポート
import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(time.Now()) // 現在時間
	fmt.Println(user.Current())
}
