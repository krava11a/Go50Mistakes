package main

import "fmt"

func main() {
	x := "text"

	//BAD
	//x[0] = 'T' //error of compile

	fmt.Println(x)

	//GOODWAY
	xbytes := []byte(x)
	xbytes[0] = 'T'
	fmt.Println(string(xbytes))

	//RUNES
	xrunes := []rune(x)
	xrunes[0] = '不'
	fmt.Println(string(xrunes))
}
