package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	sum := 0

	for ; a <= b; a++ {
		// fmt.Println(a)
		sum += a
	}
	fmt.Println(sum)
}
