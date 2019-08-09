package main

import (
	"fmt"
	"learn-go/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func main() {
	// r := mock.Retriever{Contents: "test text..."}
	r := real.Retriever{UserAgent: "", TimeOut: 3000}
	fmt.Println(download(r))
}
