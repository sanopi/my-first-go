package main

import "fmt"

func main() {
	// 1-奇数 みたいな感じで出す

	// for i := range 100 {
	// 	if (i+1)%2 == 1 {
	// 		fmt.Printf("%d-奇数\n", i+1)
	// 	} else {
	// 		fmt.Printf("%d-偶数\n", i+1)
	// 	}
	// }

	for i := 1; i <= 100; i++ {
		switch {
		case i%2 == 1:
			fmt.Printf("%d-奇数\n", i)
		case i%2 == 0:
			fmt.Printf("%d-偶数\n", i)
		}
	}

}
