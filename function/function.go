package main

func main() {
	// evenOdd()
	swapping1()
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

func swapping1() {
	n, m := swap(10, 20)
	println(n, m)
}

func swap(a int, b int) (int, int) {
	return b, a
}
