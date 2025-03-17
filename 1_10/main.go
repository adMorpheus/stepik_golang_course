package main

import "fmt"

func main() {
	sum := 0
	// Цикл от 1 до 9
	for i := 1; i < 10; i++ {
		sum += i
		fmt.Println(i)
	}
	fmt.Println(sum)
}
