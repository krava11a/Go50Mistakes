package main

import "fmt"

func main() {
	//Не всегда известно, находится ли переменная в стеке или в куче.
	//Если в С++ создать переменную с помощью оператора new, то она всегда будет в куче.
	//В Go место размещения переменной выбирает компилятор, даже если используются функции new() или make().
	//Компилятор делает выбор на основании размера и результата «анализа локальности» (escape analysis).
	//Это также означает, что можно возвращать ссылки на локальные переменные, что недопустимо в других языках,
	//например в С и С++.
	//
	//Если вы хотите знать, где находятся переменные, то передайте -gcflags -m в go build или go run
	//(например, go run -gcflags -m app.go).
	var num int
	var fp float32
	var complex complex64
	var str string
	var char rune
	var yes bool
	var events <-chan string
	var handler interface{}
	var ref *byte
	var raw [10]byte

	fmt.Println(
		num,
		fp,
		complex,
		str,
		char,
		yes,
		events,
		handler,
		ref,
		raw,
	)
}
