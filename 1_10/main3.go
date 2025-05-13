package main

import "fmt"

func main() {
	var count, sum, num int
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		fmt.Scan(&num)
		if 9 < num && num < 100 && num%8 == 0 {
			sum += num
		}
	}
	fmt.Println(sum)

}
