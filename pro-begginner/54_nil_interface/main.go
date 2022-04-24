package main

import "fmt"

func main() {
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil)
	fmt.Println(in, in == nil)

	in = data
	fmt.Println(in, in == nil)


	//IT's TRAP!!!!!
	doit := func(arg int) interface{} {
		var result *struct{} = nil

		if(arg > 0) {
			result = &struct{}{}
		}

		return result
	}

	if res := doit(-1); res != nil {
		fmt.Println("good result:",res) // выводит: good result: <nil>
		// 'res' не является 'nil', но его значение — 'nil'
	}

	//It's CORRECT
	doit2 := func(arg int) interface{} {
		var result *struct{} = nil

		if(arg > 0) {
			result = &struct{}{}
		} else {
			return nil // возвращает явный 'nil'
		}

		return result
	}

	if res := doit2(-1); res != nil {
		fmt.Println("good result:",res)
	} else {
		fmt.Println("bad result (res is nil)") // здесь — как и ожидалось
	}
}
