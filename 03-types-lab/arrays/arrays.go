package main

import "fmt"

func process(x [2]string) int {
	sum := 0
	for _, v := range x {
		sum += len(v)
	}
	return sum
}

func lengths(x *[3]string) [3]int {
	var results [3]int
	for i, v := range x {
		results[i] += len(v)
	}
	return results
}
func main() {
	var arr [2]string
	arr[0] = "Hello"
	arr[1] = "Golang"

	b := arr
	b[1] = "Java"

	for i, s := range arr {
		fmt.Printf("arr[%d] : %#s\n", i, s)
	}

	for i, s := range b {
		fmt.Printf("b[%d] : %#s\n", i, s)
	}

	fmt.Println(process(arr))
}
