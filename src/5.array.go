package main

import "fmt"

// arr is pass value!
func printArray(arr [2]float32) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [2]float32{1.0, 2.4}
	arr3 := [...]bool{true, false}

	// [2][3] row column
	arr4 := [2][3]string{
		{"1", "2"},
		{"3"},
	}
	fmt.Println(arr1, arr2, arr3, arr4)

	printArray(arr2)
}
