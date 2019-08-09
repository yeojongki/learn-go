package main

import (
	"fmt"
	"learn-go/retriever/mock"
	"learn-go/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{Contents: "test text..."}
	r = &real.Retriever{UserAgent: "", TimeOut: time.Minute}
	inspect(r)
	// fmt.Println(download(r))

	// Type assertion
	// realRetriever := r.(*real.Retriever)

	// will error
	// mockRetriever := r.(mock.Retriever)

	// fmt.Println(mockRetriever.Contents)

	// avoid panic
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("TimeOut:", v.TimeOut)
	}
}
