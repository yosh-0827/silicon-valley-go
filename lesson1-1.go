package main

import (
	"fmt"
	"os/user"
	"time"
)

// 本ファイルは主に出力（Printlnを触る）

/*
* init関数はmain関数よりも前に呼ばれる
 */
func init() {
	fmt.Println("init")
}

func bazz() {
	fmt.Println("Bazz")
}

// コメントアウトは２週類ある。「//」か「/* */」
func main() {
	bazz()
	fmt.Println("Hello World")
	fmt.Println("Hello World", "Pizza")
	fmt.Println(time.Now())
	fmt.Println(user.Current())
}
