package main

import "fmt"

func Sum(x, y int) int {
	sum := x
	for i := 1; i < y; i++ {
		sum = sum + x
	}
	return sum
}
func main() {
	fmt.Println(Sum(2, 3))
}
