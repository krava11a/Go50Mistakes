package main

import (
	"encoding/json"
	"fmt"
)

type MyData struct {
	One int
	two string
}

func main() {
	in := MyData{
		One: 1,
		two: "two",
	}
	fmt.Printf("%#v\n", in)

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded))

	var out MyData
	json.Unmarshal(encoded,&out)
	fmt.Printf("%#v\n",out)
}
