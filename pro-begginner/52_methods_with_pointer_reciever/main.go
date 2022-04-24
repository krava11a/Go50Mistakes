package main

import "fmt"

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:",p.name)
}

type printer interface {
	print()
}

func main() {
	d1 := data{"one"}
	d1.print() //ok


	////NOT ADDRESSABLE INTERFACE
	//var in printer = data{"two"} // ошибка
	//in.print()
	////NOT ADDRESSALE MAP
	//m := map[string]data {"x":data{"three"}}
	//m["x"].print() //ошибка


	//my test
	z:= map[string]printer{}
	z["x"]= &d1
	z["y"] = &data{"two"}
	for k, v := range z {
		v.print()
		fmt.Println(k)
	}
}