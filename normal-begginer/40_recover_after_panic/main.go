package main

import "fmt"

func doRecover() {
	fmt.Println("recovered =>",recover())
}

func main() {

	//not work
	//recover() //nothing to do
	//panic("not good")
	//recover()// not work
	//fmt.Println("ok")

	//better
	//defer func() {
	//	fmt.Println("recovered:",recover())
	//}()
	//
	//panic("not good, but better")

	//not work
	defer func() {
		doRecover()
	}()

	panic("not good")

}
