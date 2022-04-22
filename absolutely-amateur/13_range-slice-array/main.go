package main

import "fmt"

func main() {

	x:=[]string{"aa","bb","vv"}

	//wrong
	for v := range x {
		fmt.Println(v)
	}
	//right
	for _, v := range x {
		fmt.Println(v)
	}
	
}
