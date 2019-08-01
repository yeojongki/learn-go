package main

import "fmt"

// out vars
var a = 1

// error -> b:=2

func varsWithOut() {
	fmt.Println(a)
}

func varsWithoutInit() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func varsWithInit() {
	var a, b int = 1, 2
	var s string = "string"
	fmt.Println(a, b, s)
}

func varsTypeDeduction() {
	var a, b, c = 1, true, "sss"
	fmt.Println(a, b, c)
}

func varsShorter() {
	a, b, c := 1, true, "ccc"
	fmt.Println(a, b, c)
}

func main() {
	varsWithOut()
	varsWithoutInit()
	varsWithInit()
	varsTypeDeduction()
	varsShorter()
}
