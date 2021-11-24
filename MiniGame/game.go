package main

import "fmt"

func removeElem(s []int, index int) []int {
	length := len(s) - 1
	for i := index; i < length; i++ {
		s[i] = s[i+1]
	}
	return s[:length]
}
func generateArray(size int) []int {

	var sli []int
	for i := 1; i <= size; i++ {
		sli = append(sli, i)
	}
	return sli
}
func FindWinner(n, m int) int {

	slice := generateArray(n)
	removeAtIndex := m - 1
	valueToMove := m - 1

	for len(slice) > 1 {

		removeAtIndex = removeAtIndex % len(slice)
		slice = removeElem(slice, removeAtIndex)
		removeAtIndex = removeAtIndex + valueToMove

	}
	return slice[0]
}
func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	result := FindWinner(n, m)
	fmt.Printf("%d", result)
}
