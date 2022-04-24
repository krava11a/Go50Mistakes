package main

import "fmt"

type data struct {
	name string
}

func main() {
	//BAD MAPS
	//m := map[string]data{"x": {"one"}}
	//m["x"].name = "two" //cannot assign to struct field m["x"].name in map

	//GOOD SLICES ))
	s := []data{{name: "oneSlice"}}
	s[0].name = "twoSlice"
	fmt.Println(s)

	//Variant with MAPS
	//1  temp variable
	m1 := map[string]data {"x": {"one"}}
	r := m1["x"]
	r.name = "two"
	m1["x"] = r
	fmt.Printf("%v\n", m1) //выводит: map[x:{two}]

	//2  map with pointers
	m2 := map[string]*data{"x":{"pointer1"}}
	m2["x"].name = "pointerNext"
	fmt.Println(m2["x"])

	//?    invalid memory address or nil pointer dereference
	m3 := map[string]*data{"x":{"pointer1"}}
	m3["z"].name = "WHAT"
	fmt.Println(m3["z"])


}
