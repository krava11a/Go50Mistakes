package main

import (
	"fmt"
	"runtime"
)

func main() {

	//1
	//done :=false
	//go func() {
	//	done = true
	//}()
	//v := 0
	//for !done {
	//	v++
	//	fmt.Println(v,"not done!") // не встроена
	//}
	//fmt.Println("done!!!")

	//2
	done := false

	go func(){
		done = true
	}()

	for !done {
		fmt.Println("not done!")
		runtime.Gosched()
	}
	fmt.Println("done!")
}
