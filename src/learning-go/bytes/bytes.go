package main

import "fmt"
import "bytes"

func main() {
	var foo = []byte("foo")
	var bar = []byte("bar")

	var result = bytes.Join([][]byte{foo, bar}, []byte{})
	fmt.Printf("%s", result)
}
