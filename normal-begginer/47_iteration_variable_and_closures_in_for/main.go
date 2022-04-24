package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (f *field) print() {
	fmt.Println(f.name)
}

func main() {

	////1
	//data := []string{"one", "two", "three"}
	//for idGorutine, v := range data {
	//	vc := v
	//	go func() {
	//		fmt.Println(idGorutine, vc)
	//	}()
	//}
	//
	//time.Sleep(3 * time.Second)
	//
	////2
	//
	//for idGorutine, v := range data {
	//
	//	go func(val string) {
	//		fmt.Println(idGorutine, val)
	//	}(v)
	//}
	//
	//time.Sleep(3 * time.Second)
	//
	////Not correct
	//data3 := []field{{"one"}, {"two"}, {"three"}}
	//
	//for _, v := range data3 {
	//	go v.print()
	//}
	//time.Sleep(3 * time.Second)
	//
	////Correct
	//data4 := []field{{"one"}, {"two"}, {"three"}}
	//time.Sleep(2*time.Second)
	//for _, v := range data4 {
	//	v := v
	//	go v.print()
	//}
	//time.Sleep(3 * time.Second)

	//?
	data5 := []*field{{"one"},{"two"},{"three"}}
	for _,v := range data5 {
		go v.print()
	}
	time.Sleep(3 * time.Second)

}
