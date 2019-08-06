package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func converToBin(num int) string {
	if num == 0 {
		return "0"
	}
	result := ""
	for ; num > 0; num /= 2 {
		left := num % 2
		result = strconv.Itoa(left) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func while() {
	for {
		fmt.Println("while")
	}
}

func main() {
	fmt.Println(
		converToBin(0),
		converToBin(2),
		converToBin(5),
		converToBin(13),
	)

	printFile("2.test.txt")
	// while()
}
