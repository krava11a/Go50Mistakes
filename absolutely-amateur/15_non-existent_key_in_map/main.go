package main

import "fmt"

func main() {

	//BAD
	x := map[string]string{"aa": "a", "two": "", "three": "c"}

	if v := x["two"]; v == "" {
		fmt.Println("BAD WAY: no entry")
	}

	//GOOD
	if v, ok := x["two"]; !ok {
		fmt.Println("GOODSPEED: no entry")
	} else {
		fmt.Println("Value of element: ", v)
	}

}
