package main

import (
	"bytes"
	"fmt"
)

func main() {

	var w Writer = ConsoleWriter{}
	w.Write([]byte ("Hello GO!"))
	fmt.Println("")

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte ("Hello Youtube Listeners, This is a test run!"))
	wc.Close()
}

/**
Declare Interface
*/
type Writer interface { // Write + er (method name)
	// describe behaviors
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

/**
Composed Interface
*/
type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, nil
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
