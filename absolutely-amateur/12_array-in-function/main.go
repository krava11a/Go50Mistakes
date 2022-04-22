package main

import "fmt"

func main() {

	//array
	fmt.Println("arrays")
	x:= [3]int{1,2,3}

	//wrong
	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)

	fmt.Println(x)

	//right
	func(arr *[3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(&x)

	fmt.Println(x)

	//slices
	fmt.Println("slices")
	s := []int{1,2,3}

	func(k []int){
		k[1] = 23
		fmt.Println(k)
	}(s)
	fmt.Println(s)
}
