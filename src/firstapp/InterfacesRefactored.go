package main

import "fmt"

func main() {
	var wc WriterCloser = &myWriterCloser{}
	fmt.Println(wc)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type myWriterCloser struct{}

func (mwc myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

func (mwc myWriterCloser) Close() error {
	return nil
}