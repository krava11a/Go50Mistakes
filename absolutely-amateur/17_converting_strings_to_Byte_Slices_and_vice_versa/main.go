package main

import (
	"fmt"
	"reflect"
)

func main() {

	t := "test"
	b := []byte(t)
	fmt.Println(t, &t)
	fmt.Println(b, &b)

	//FIRST OPTIMIZATION
	m := make(map[string]string)
	m["p"] = "1"
	m["str"] = "2"
	t = "p"
	b = []byte(t)
	if reflect.DeepEqual(m[string(b)], m["p"]) {
		fmt.Println("First OPTIMIZATION")
	}

	//SECOND OPTIMIZATION
	s := "FIRST OPTIMIZATION"
	for i, v := range []byte(s) {
		fmt.Println("position ", i)
		fmt.Println("value ", v, string(v))
	}

}
