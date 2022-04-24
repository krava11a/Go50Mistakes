package main

import "fmt"

func main() {
	data := []int{1, 2, 3}
	for _, v := range data {
		v *= 10
	}
	fmt.Println("NOT CHANGED data:", data)

	for i, _ := range data {
		data[i] *= 10
	}
	fmt.Println("CHANGED data:", data)

	//BUT IF SLICE HAS POINTERS VALUE
	dataPointers := []*struct{ num int }{{1}, {2}, {3}}

	for _,v := range dataPointers{
		v.num *=10
	}

	fmt.Println(dataPointers[0],dataPointers[1],dataPointers[2],)
}
