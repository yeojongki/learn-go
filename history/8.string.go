package main

import (
	"fmt"
	"unicode/utf8"
)

// any inner type can be a key excluding `slice, map, function`
func main() {
	text := "abc测试."

	// 61 62 63 E6 B5 8B E8 AF 95 2E
	// 每个中文 3 字节
	for _, v := range []byte(text) {
		// 16 进制
		fmt.Printf("%X ", v)
	}
	fmt.Println()

	// utf8 -> unicode
	// (0 61) (1 62) (2 63) (3 6D4B) (6 8BD5) (9 2E)
	// `测` 为 4 字节整数 int32 相当于一个 rune
	for i, v := range text {
		fmt.Printf("(%d %X) ", i, v)
	}
	fmt.Println()

	// text length
	fmt.Println("text length:", utf8.RuneCountInString(text))

	bytes := []byte(text)
	// fmt.Println(utf8.DecodeRune(bytes)) // 97 1

	for len(bytes) > 0 {
		char, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		// %c 输出定整数的对应字符 61 对应 a
		fmt.Printf("%c", char)
	}
	fmt.Println()

	// 转 rune
	for i, char := range []rune(text) {
		fmt.Printf("(%d %c )", i, char)
	}
}
