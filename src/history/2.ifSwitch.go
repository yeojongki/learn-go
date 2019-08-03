package main

import (
	"fmt"
	"io/ioutil"
)

const filename = "2.test.txt"

func if1st() {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		// if no file will log: open 2.test.txt: The system cannot find the file specified.
		fmt.Println(err)
	} else {
		fmt.Printf("%s", content)
	}
}

func if2nd() {
	if content, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}
	// if 条件赋值的变量作用域只在这个 if 语句里
	// error
	// fmt.Printf(content)
}

func switchAutoBreak(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "C"
	case score < 80:
		g = "B"
	case score < 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Error score: %d", score))
	}
	return g
}

func main() {
	if2nd()
	fmt.Println(
		switchAutoBreak(59),
		switchAutoBreak(79),
		switchAutoBreak(99),
		switchAutoBreak(199),
	)
}
