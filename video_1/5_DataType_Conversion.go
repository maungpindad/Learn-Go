package main

import "fmt"

func main() {
	var value32 int32 = 32769
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value32) // Buffer overflow

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value16)

	var manchester = "red"
	var d = manchester[2]
	var dString = string(d)

	fmt.Println(manchester)
	fmt.Println(d)
	fmt.Println(dString)
}
