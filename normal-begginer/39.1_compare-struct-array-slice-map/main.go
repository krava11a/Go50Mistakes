package main

import (
	"fmt"
	"reflect"
)

//compare
type data struct {
	num     int
	fp      float32
	complex complex64
	str     string
	char    rune
	yes     bool
	events  <-chan string
	handler interface{}
	ref     *byte
	raw     [10]byte
}

//not compare
type dataNotCompare struct {
	num    int               //ok
	checks [10]func() bool   //not comparable
	doit   func() bool       //not comparable
	m      map[string]string //not comparable
	bytes  []byte            //not comparable
}

func (d *dataNotCompare) getNum() int {
	return d.num
}

func main() {

	v1 := data{}
	v2 := data{}
	fmt.Println("v1==v2:", v1 == v2)

	v3 := dataNotCompare{

		doit: func() bool {
			return true
		},
	}
	v4 := dataNotCompare{
		doit: func() bool {
			return false
		},
	}
	fmt.Println("v3==v4:", reflect.DeepEqual(v3, v4))
	v3.num = 3
	v4.num = 4

	fmt.Println("v3.num =", v3.getNum(), "v4.num =", v4.getNum())
	fmt.Println("v3.doit =", v3.doit(), "v4.doit =", v4.doit())

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2)) // prints: m1 == m2: true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2)) // prints: s1 == s2: true



}
