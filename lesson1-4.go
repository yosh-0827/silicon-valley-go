// 本ファイルはデータ構造を扱う
package main

import (
	"fmt"
)

func main() {
	// 配列
	var a [2]int // int型の要素を２つ持つ配列を作成
	a[0] = 100   // 1個目の要素に入れる
	a[1] = 200   // ２個目の要素に入れる
	fmt.Println(a)

	b := [2]int{100, 200} // 上記を短縮して書くことも可能
	fmt.Println(b)

	// スライス(要素数を持たない配列みたいなやつ)
	n := []int{1, 2, 3, 4, 5}
	n[2] = 100 // ３個目の要素を書き換えることも可能
	fmt.Println(n, n[2])
	fmt.Println(n[2:4]) // 特定の範囲の要素を出力
	fmt.Println(n[:2])
	fmt.Println(n[1:])
	fmt.Println(n[:]) // 全ての要素を出力

	// 配列に要素を追加する(配列には使えない。スライスのみ可能)
	n = append(n, 100)
	fmt.Println(n)

	// 多次元スライス（np.arrayみたいなやつ）
	// 3行3列の行列
	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)
	fmt.Println(board[1])    // ２個目の行を出力
	fmt.Println(board[1][2]) // ２個目の行の３個目の要素を出力

	// make関数で要素が０で初期化されたスライスを作成できる
	n_make := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v", len(n_make), cap(n), n_make) // lenは要素数, capは容量

	// map: map[キーの型]値の型{値の辞書}
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	m["banana"] = 300 // 値の書き換え
	fmt.Println(m)
	m["orange"] = 500 // 新しい要素を追加
	fmt.Println(m)
	v, ok := m["apple"]
	fmt.Println(v, ok) // 要素が存在するかのチェック
	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)
	m_non := map[string]int{} // 空のMapを作ってそこに値を入れることもできる
	m_non["pc"] = 5000
	fmt.Println(m_non)

}
