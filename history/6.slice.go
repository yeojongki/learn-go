package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d \n", s, len(s), cap(s))
}

func createSlice() {
	s6 := []int{1, 2, 3}
	printSlice(s6)

	s7 := make([]int, 10)
	s8 := make([]int, 10, 20)
	printSlice(s7)
	printSlice(s8)
}

func copySlice() {
	s9 := []int{1, 2}
	s10 := make([]int, 10, 20)
	copy(s10, s9)
	printSlice(s10)
}

func deleteSlice() {
	s11 := []int{1, 3, 5, 7, 9, 11}
	// delete 7
	s11 = append(s11[:3], s11[4:]...)
	printSlice(s11)

	// shift
	head := s11[0]
	s11 = s11[1:]

	// pop
	tail := s11[len(s11)-1]
	s11 = s11[:len(s11)-1]
	printSlice(s11)
	fmt.Println(head, tail)
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// s1 := arr[:]
	// s2 := arr[0:2]
	// fmt.Println(s1, s2)

	fmt.Printf("before arr=%v \n", arr)

	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d \n", s1, len(s1), cap(s1)) // s1=[2 3 4 5], len(s1)=4, cap(s1)=6
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d \n", s2, len(s2), cap(s2)) // s2=[5 6], len(s2)=2, cap(s2)=3

	s3 := append(s2, 10)
	s4 := append(s3, 20)
	s5 := append(s4, 30)
	fmt.Printf("s3=%v, s4=%v, s5=%v \n", s3, s4, s5) // s3=[5 6 10], s4=[5 6 10 20], s5=[5 6 10 20 30]

	fmt.Printf("after arr=%v, len(arr)=%d \n", arr, len(arr)) // [0 1 2 3 4 5 6 10], len(arr)=8

	createSlice()

	copySlice()

	deleteSlice()
}
