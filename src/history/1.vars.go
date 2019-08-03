package main

import (
	"fmt"
	"math"
)

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

func forceTransform() {
	a, b := 3, 4
	c := math.Sqrt(float64(a*a + b*b))
	fmt.Println(c)
}

func consts() {
	// const 数值可以作为各种类型使用  对比函数 `forceTransform`
	const a, b = 3, 4
	c := math.Sqrt(a*a + b*b)
	fmt.Println(c)
}

func enums() {
	const (
		a = iota
		_
		c
		d
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(a, b, c, d)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	varsWithOut()
	varsWithoutInit()
	varsWithInit()
	varsTypeDeduction()
	varsShorter()

	forceTransform()
	consts()
	enums()
}
