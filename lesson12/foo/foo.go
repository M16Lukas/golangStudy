package foo

const (
	Max = 100 // 一文字が大文字 = パブリック、外部のパッケージから参照可能
	min = 1   // 一文字が小文字 = プライベート、外部のパッケージから参照不可
)

func ReturnMin() int {
	return min
}
