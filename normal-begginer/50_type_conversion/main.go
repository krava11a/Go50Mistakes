package main

import "fmt"

func main() {

	//NOT CORRECT
	var data interface{} = "great"

	if data, ok := data.(int); ok {
		fmt.Println("[is an int] value =>", data)
	} else {
		fmt.Println("[not an int] value =>", data)
		//выводит: [not an int] value => 0 (not "great")
	}

	//CORRECT

	if res, ok := data.(int); ok {
		fmt.Println("[is an int] value =>",res)
	} else {
		fmt.Println("[not an int] value =>",data)
		// выводит: [not an int] value => great (as expected)
	}
}
