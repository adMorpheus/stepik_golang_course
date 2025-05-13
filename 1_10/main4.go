package main

import "fmt"

func main() {
	var num, count, max int

	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		if num > max {
			max = num
			count = 1
			fmt.Println("more then")
			fmt.Println("count: ", count)

		} else if num == max {
			fmt.Println("equal")
			count++
			fmt.Println("count: ", count)
		}
	}

	fmt.Print(count)

}
