package main

import "fmt"

const hello = "hello"

func main() {
	world := "world"
	fmt.Println(hello + ", " + world + "!")
	fmt.Printf("%c", 'a'+10) // 0-indexedで10番目のアルファベットが出力される。
}
