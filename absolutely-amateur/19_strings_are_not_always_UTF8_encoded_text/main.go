package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	data1 := "ABC"
	fmt.Println(utf8.ValidString(data1))

	data2 :="a\xfec"
	fmt.Println(utf8.ValidString(data2))
}
