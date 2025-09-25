package main

func main() {
	evenOdd()
}

func evenOdd() {
	for i := 1; i < 100; i++ {
		print(i)
		if isEven(i) {
			println("-偶数")
		} else {
			println("-奇数")
		}
	}
}

func isEven(i int) bool {
	return i%2 == 0
}
