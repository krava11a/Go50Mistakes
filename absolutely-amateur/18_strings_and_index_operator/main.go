package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "text"
	fmt.Println(x[0])
	fmt.Printf("%T\n", x[0])
	fmt.Println(string(x[0]))
	x = strings.Replace(x,"t","x",2)
	//for _, v := range x{
	//	fmt.Println(string(v))
	//}
	fmt.Println(x)
}
