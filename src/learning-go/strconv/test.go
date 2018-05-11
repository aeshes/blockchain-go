package main

import "strconv"
import "fmt"

func main() {
	var timestamp int64 = 123456
	conv := strconv.FormatInt(timestamp, 2)

	fmt.Println(conv)

	conv2 :=[]byte(conv)
	fmt.Println(conv2)
}
