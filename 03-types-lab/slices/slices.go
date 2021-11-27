package main

import "fmt"

func printInterfaceSlice(message string, x []interface{}) {
	fmt.Printf("%s  len=%d  cap= %d  %#v\n", message, len(x), cap(x), x)
}

func printSlice(message string, x []string) {
	fmt.Printf("%s  len=%d  cap= %d  %#v\n", message, len(x), cap(x), x)
}
func printIntSlice(message string, x []int) {
	fmt.Printf("%s  len=%d  cap= %d  %#v\n", message, len(x), cap(x), x)
}
func main() {
	a := [...]string{"Hello", "Golang", "!"}

	sa1 := a[:]
	printSlice("sa1 is", sa1)

	b := make([]int, 5, 10)
	printIntSlice("b is", b)

	c := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printIntSlice("c is ", c)
	sc1 := c[2:5:10]
	printIntSlice("sc1 is ", sc1)
	sc2 := sc1[2:3:3]
	printIntSlice("sc2 is ", sc2)

	sc3 := append(sc1, 11, 12)
	printIntSlice("sc1 is ", c)
	printIntSlice("sc3 is ", sc3)

	var sc4 []interface{}

	for _, v := range c {
		sc4 = append(sc4, v)
	}
	printInterfaceSlice("sc4 is", sc4)

	//var sc5 []int
	sc5 := make([]int, 10)
	copy(sc5, c)
	printIntSlice("sc5 is", sc5)
}
