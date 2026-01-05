// 主にデータ型を扱う
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Goは１番長い部分位合わせてかくので、この場合はc64にインデントを揃える
var (
	u8  uint8     = 255 // 正の値のみをとる型で8bitまで。uint16は16bitまで扱える
	i8  int8      = 127
	f32 float32   = 0.2      // float
	c64 complex64 = -5 + 12i // 複素数
)

func main() {
	fmt.Println(u8, i8, f32, c64)
	fmt.Printf("type=%T value=%v\n", u8, u8) // %Tは型を表示, %vは値を表示, %tはboolを表示, %sは文字列
	fmt.Println("hello ")

	// 文字列連結
	fmt.Println("Hello" + "World")
	// インデックス(例として０番目のHがASCIIで出る))
	fmt.Println("Hello"[0])
	// インデックスの文字を出力
	fmt.Println(string("Hello"[0]))

	//　文字列の置き換え
	var s string = "Hello World"
	s = strings.Replace(s, "H", "X", 1) // 第一：対象、第二：置き換え対象、第３：置き換え後の文字、第四：何回置き換えるか
	fmt.Println(s)

	// キャスト
	var s2 string = "14"
	i, _ := strconv.Atoi(s2) // stringからintはStrconvを使用。実行結果とエラー時の結果を合わせて返すので変数が２ついる
	fmt.Printf("%T %v", i, i)
}
