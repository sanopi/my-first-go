package main

import "fmt"

const hello = "hello"

func main() {
	world := "world"
	fmt.Println(hello + ", " + world + "!")
	fmt.Printf("%c\n", 'a'+10) // 0-indexedで10番目のアルファベットが出力される。

	fmt.Println("1 & 2")
	fmt.Println(1 & 2)
	fmt.Println("1 &^ 2")
	fmt.Println(1 &^ 2)
	fmt.Println("2 &^ 1")
	fmt.Println(2 &^ 1)
	fmt.Println("1 &^ 1")
	fmt.Println(1 &^ 1)
	fmt.Println("0b111111 &^ 0b101010")
	fmt.Println(0b111111 &^ 0b101010) // 0b010101 -> 16+4+1 = 21

}
