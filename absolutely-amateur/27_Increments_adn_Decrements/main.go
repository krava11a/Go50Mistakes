package main

import "fmt"

func main() {

	//don't use increment/decrement as a preffix
	data := []int{1, 2, 3}
	i:=0
	i++
	//fmt.Println(data[i++]) //mistake
	fmt.Println(data[i])
}
