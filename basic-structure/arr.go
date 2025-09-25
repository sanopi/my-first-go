package main

import "fmt"

func main() {
	var arr [10]int
	for i := range 10 {
		arr[i] = i
	}

	// sl := arr[2:8]
	// fmt.Println(sl)
	// sl1 := append(sl, 80, 90, 100)
	// fmt.Println(sl)
	// fmt.Println(sl1)

	sll := arr[:3:8]
	fmt.Println(sll)
	fmt.Println(append(sll, 100))
	fmt.Println(arr) // => [0 1 2 100 4 5 6 7 8 9]
	fmt.Println(sll)
}
