package main

import "fmt"

var num int

func main() {
	for {
		fmt.Scan(&num)
		if num < 10 {
			continue
		} else if num > 100 {
			break
		} else {
			fmt.Println(num)
		}

	}
}
