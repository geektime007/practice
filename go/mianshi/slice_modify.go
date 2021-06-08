package main

import "fmt"

func sliceModify(slice []int) {
	slice[0] = 88
	slice = append(slice, 6)
	fmt.Println("==", slice)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	sliceModify(slice)
	fmt.Println(slice)
}
