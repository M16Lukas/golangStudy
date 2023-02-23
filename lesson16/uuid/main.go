package main

import (
	"fmt"

	"github.com/google/uuid"
)

/*
------------------------------------------------------------------
* UUID（Universally Unique Identifier）
* ソフトウェア上でオブジェクトを一意に識別するための識別子
------------------------------------------------------------------
*/

func main() {
	// NewUUID
	uuidObj, _ := uuid.NewUUID()
	fmt.Println("  ", uuidObj.String()) //   de6b207a-b37b-11ed-bc3f-9eb6d0c51433

	// NewMD5
	data := []byte("wnw8olzvmjp0x6j7ur8vafs4jltjabi0")
	uuidObj2 := uuid.NewMD5(uuidObj, data)
	fmt.Println("  ", uuidObj2.String()) //    194babb4-d862-3449-842f-6988a361a9a9

	// NewSHA1
	data2 := []byte("wnw8olzvmjp0x6j7ur8vafs4jltjabi0")
	uuidObj3 := uuid.NewSHA1(uuidObj, data2)
	fmt.Println("  ", uuidObj3.String()) //    83aba0d2-88e1-5915-9a67-911ff90863c6

	// NewRandom
	uuidObj4, _ := uuid.NewRandom()
	fmt.Println("  ", uuidObj4.String()) //   71b8977e-07de-42b5-a2ac-7d4967229323
}
