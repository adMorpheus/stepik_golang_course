package main

import (
	"fmt"
	"math"
)

var x, p, y float64

func main() {
	fmt.Scan(&x, &p, &y)
	p = p/100.0 + 1
	fmt.Println(p)
	for i := 0; ; i++ {
		if x >= y {
			fmt.Println(i)
			break
		}
		x = x * p
		math.Round(x)
		fmt.Println(x)
	}
}
