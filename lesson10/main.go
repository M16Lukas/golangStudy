package main

import "fmt"

/*
-------------------------------------------------------------
* 構造体
* オブジェクト志向のプログラミング言語のclassのような存在
* 複数の任意の型の値を一つにまとめたもの
-------------------------------------------------------------
*/

/*
-------------------------------------------------------------
* struct
-------------------------------------------------------------
*/
type User struct {
	// フィールド
	Name string
	Age  int
	// X, Y int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

/*
-------------------------------------------------------------
* struct メソッド
-------------------------------------------------------------
*/
func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

// 構造体に定義するメソッドのレシーバーは基本的にはポインタ型にすべき
func (u *User) SetNameV2(name string) {
	u.Name = name
}

/*
-------------------------------------------------------------
* struct 埋め込み
-------------------------------------------------------------
*/
type T struct {
	User
}

func (u *User) SetNameV3() {
	u.Name = "C"
}

/*
-------------------------------------------------------------
* struct 型コンストラクタ
-------------------------------------------------------------
*/
func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

/*
-------------------------------------------------------------
* struct スライス
-------------------------------------------------------------
*/
type Users []*User

/*
-------------------------------------------------------------
* struct 独自型
-------------------------------------------------------------
*/
type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	/* -------------------------------------------------------------
	* struct
	-------------------------------------------------------------*/
	var user1 User
	fmt.Println(user1) // { 0}
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1) // {user1 10}

	user2 := User{}
	fmt.Println(user2) // { 0}
	user2.Name = "user2"
	fmt.Println(user2) // {user2 0}

	user3 := User{
		Name: "user3",
		Age:  30,
	}
	fmt.Println(user3) // {user3 30}

	user4 := User{"user4", 40}
	fmt.Println(user4) // {user4 40}

	user5 := User{Name: "user5"}
	fmt.Println(user5) // {user5 0}

	user6 := new(User) // 構造体のポインタ型を返す
	fmt.Println(user6) // &{ 0}

	// アドレス演算子をつけて宣言することが多い
	// 関数の引数として、構造体の変数を渡す場合にポインタ型をよく使う
	user7 := &User{}
	fmt.Println(user7) // &{ 0}

	UpdateUser(user1)
	UpdateUser2(user7)

	fmt.Println(user1) // {user1 10}
	fmt.Println(user7) // &{A 1000}

	/* -------------------------------------------------------------
	* struct メソッド
	-------------------------------------------------------------*/
	user8 := User{Name: "user8"}
	user8.SayName() // user8

	user8.SetName("A")
	user8.SayName() // user8

	user8.SetNameV2("A")
	user8.SayName() // A

	user9 := &User{Name: "user9"}
	user9.SetNameV2("B")
	user9.SayName() // B

	/* -------------------------------------------------------------
	* struct 埋め込み
	-------------------------------------------------------------*/
	t := T{
		User: User{Name: "user10", Age: 10},
	}
	fmt.Println(t)           // {{user10 10}}
	fmt.Println(t.User)      // {user10 10}
	fmt.Println(t.User.Name) // user10
	fmt.Println(t.Name)      // user10

	t.User.SetNameV3()
	fmt.Println(t) // {{C 10}}

	/* -------------------------------------------------------------
	* struct 型コンストラクタ
	-------------------------------------------------------------*/
	user11 := NewUser("user11", 40)
	fmt.Println(user11)  // &{user11 40}
	fmt.Println(*user11) // {user11 40}

	/* -------------------------------------------------------------
	* struct スライス
	-------------------------------------------------------------*/
	user12 := User{"user12", 23}
	user13 := User{"user13", 33}
	user14 := User{"user14", 43}
	user15 := User{"user15", 53}
	user16 := User{"user16", 63}

	users := Users{}
	users = append(users, &user12, &user13, &user14, &user15, &user16)

	fmt.Println(users) // [0xc000008240 0xc000008258 0xc000008270 0xc000008288 0xc0000082a0]
	for _, u := range users {
		fmt.Println(*u)
	}
	/*
		{user12 23}
		{user13 33}
		{user14 43}
		{user15 53}
		{user16 63}
	*/

	users2 := make([]*User, 0)
	users2 = append(users2, &user12, &user13)
	for _, u := range users2 {
		fmt.Println(*u)
	}
	/*
		{user12 23}
		{user13 33}
	*/

	/* -------------------------------------------------------------
	* struct マップ
	-------------------------------------------------------------*/
	m := map[int]User{
		1: {Name: "user17", Age: 34},
		2: {Name: "user18", Age: 64},
	}
	fmt.Println(m) // map[1:{user17 34} 2:{user18 64}]

	m2 := map[User]string{
		{Name: "user18", Age: 84}: "tokyo",
		{Name: "user19", Age: 14}: "LA",
	}
	fmt.Println(m2) // map[{user18 84}:tokyo {user19 14}:LA]

	m3 := make(map[int]User)
	fmt.Println(m3) // map[]
	m3[1] = User{"user20", 19}
	fmt.Println(m3) // map[1:{user20 19}]

	for _, v := range m {
		fmt.Println(v)
	}
	/*
		{user17 34}
		{user18 64}
	*/

	/* -------------------------------------------------------------
	* struct 独自型
	-------------------------------------------------------------*/
	var mi MyInt
	fmt.Println(mi)        // 0
	fmt.Printf("%T\n", mi) // main.MyInt
	mi.Print()             // 0

	// i := 100
	// fmt.Println(i + mi) // mismatched types int and MyInt)
}
