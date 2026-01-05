// 本ファイルは変数を扱う

package main

import (
	"fmt"
)

// 定数(グローバル変数)はconstで宣言
const Pi = 3.14
const (
	UserName = "test_user"
	Password = "test_pass"
)

// 変数のスコープとして、varは関数外でもOK。短縮宣言は関数内のみ。

func main() {
	// var 変数名 型 = 値　で変数宣言
	var i int = 1
	var t, f bool = true, false
	var (
		j      int     = 1
		f64    float64 = 1.2
		ss     string  = "test"
		tt, ff bool    = true, false
	)
	fmt.Println(i, t, f)
	fmt.Println(j, f64, ss, tt, ff)

	// 変数名 := 値でもvarを省略して変数宣言可能(型が自動定義されるため注意)
	xi := 1
	fmt.Println("短縮変数宣言", xi)
	fmt.Printf("%T\n", xi) // printfで型を表示, %Tもセットで使う。　\nで改行

	fmt.Println(Pi, UserName, Password)
}
