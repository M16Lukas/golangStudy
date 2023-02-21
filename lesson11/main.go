package main

import "fmt"

/*
--------------------------------------------------
* interface　異なる型に共通の性質を付与する
--------------------------------------------------
*/

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

/*
--------------------------------------------------
* interface　カスタムエラー
--------------------------------------------------
*/
type MyError struct {
	Message   string
	ErrorCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrorCode: 404}
}

/*
--------------------------------------------------
* interface　Stringer
--------------------------------------------------
*/
// type Stringer interface {
// 	String() string
// }
type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {

	/* --------------------------------------------------
	* interface　異なる型に共通の性質を付与する
	--------------------------------------------------*/
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "123-456", Model: "AB-1234"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}
	// Name=Taro, Age=21
	// Number=123-456, Model=AB-1234

	/* --------------------------------------------------
	* interface　カスタムエラー
	--------------------------------------------------*/
	err := RaiseError()
	fmt.Println(err.Error()) // カスタムエラーが発生しました
	// fmt.Println(err.Message) // undefined (type error has no field or method Message)

	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrorCode)
	} // 404

	/* --------------------------------------------------
	* interface　Stringer
	--------------------------------------------------*/
	p := &Point{100, "ABC"}
	fmt.Println(p) // <<100, ABC>>
}
