package main

import (
	"fmt"
	"math"
)

func main() {
	/* -----------------------------------------------------------
	 * math
	 -----------------------------------------------------------*/
	// 数学敵な定数
	// 円周率
	fmt.Println(math.Pi) // 3.141592653589793

	// 2の平方根
	fmt.Println(math.Sqrt2) // 1.4142135623730951

	// 数値型に関数る定数
	// float32で表現可能な最大値
	fmt.Println(math.MaxFloat32) // 3.4028234663852886e+38
	// float32で表現可能な０ではない最小値
	fmt.Println(math.SmallestNonzeroFloat32) // 1.401298464324817e-45
	// float64で表現可能な最大値
	fmt.Println(math.MaxFloat64) // 1.7976931348623157e+308
	// float64で表現可能な０ではない最小値
	fmt.Println(math.SmallestNonzeroFloat64) // 5e-324
	// int64 version
	fmt.Println(math.MaxInt64) // 9223372036854775807
	fmt.Println(math.MinInt64) // -9223372036854775808

	// 絶対値
	fmt.Println(math.Abs(100))  // 100
	fmt.Println(math.Abs(-100)) // 100

	// 累乗を求める
	fmt.Println(math.Pow(0, 2)) // 0
	fmt.Println(math.Pow(2, 2)) // 4

	// 平方根、立方根
	fmt.Println(math.Sqrt(2)) // 1.4142135623730951
	fmt.Println(math.Cbrt(8)) // 2

	// 最大値、最小値
	fmt.Println(math.Max(1, 2)) // 2
	fmt.Println(math.Min(1, 2)) // 1

	// 小数点以下の切り捨て、切り上げ

	// math.Trunc : 数値の正負を問わず、小数点以下を単純に切り捨てる
	fmt.Println(math.Trunc(3.14))  // 3
	fmt.Println(math.Trunc(-3.14)) // -3

	// math.Floor : 引数の数値より小さい最大の整数を返す
	fmt.Println(math.Floor(3.5))  // 3
	fmt.Println(math.Floor(-3.5)) // -4

	// math.Ceil : 引数の数値より大きい最小の整数を返す
	fmt.Println(math.Ceil(3.5))  // 4
	fmt.Println(math.Ceil(-3.5)) // -3

	n := math.Sqrt(-1)
	fmt.Println(math.IsNaN(n)) // true
}
