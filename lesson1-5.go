/*
* 本ファイルは関数を扱う
 */

package main

import (
	"fmt"
)

// 特定の文字を表示するだけの関数
func add() {
	fmt.Println("add function")
}

// 引数あり
func add2(x int, y int) {
	fmt.Println("x:", x, "y:", y)
	fmt.Println(x + y)
}

// returnあり
// func 関数名(引数名 型) 返り値の型{処理}
func add3_return(x int, y int) int {
	return x + y
}

// 返り値が２つあるケース
func add3_return2(x int, y int) (int, int) {
	return x + y, x - y
}

func main() {
	add()
	add2(10, 20)
	add3 := add3_return(100, 200)
	fmt.Println(add3)
	add3_1, add3_2 := add3_return2(10, 20) // 返り値が２つある場合は、二つ変数の受け皿が必要
	fmt.Println(add3_1, add3_2)
}
