package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func zeroValue(value int) {
	value = 0
}

func zeroPtr(prt *int) {
	*prt = 100
}

func main() {
	m := map[string]int{
		"apple":  100,
		"banana": 200,
	}

	fmt.Println(m)
	fmt.Println(len(m))
}
