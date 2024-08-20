package main

import (
	"fmt"
)

func main() {
	var a bool = true
	fmt.Println(a)

	var b int = 0
	fmt.Println(b)

	var c int8 = 127
	fmt.Println(c)

	var d uint8 = 255
	fmt.Println(d)

	var e byte = d
	fmt.Println(e)

	var f int16 = 32767
	fmt.Println(f)

	var g uint16 = 65535
	fmt.Println(g)

	var h int32 = 2147483647
	fmt.Println(h)

	var i rune = 2147483647
	fmt.Println(i)

	var j uint32 = 4294967295
	fmt.Println(j)

	var k int64 = 9223372036854775807
	fmt.Println(k)

	var l uint64 = 18446744073709551615
	fmt.Println(l)
}
