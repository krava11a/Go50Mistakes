package main

import "fmt"

func main() {

	x := map[string]string{"aa": "a", "two": "", "three": "c", "four": "captain", "five": "Hockey"}

	for k, v := range x {
		fmt.Println("key:", k, "value:", v)
	}
	
}
