package main

import "fmt"

func getBad() []byte {
	raw := make([]byte,10000)
	fmt.Println(len(raw),cap(raw),&raw[0]) // выводит: 10000 10000 <byte_addr_x>
	return raw[:3]
}

func getGood() []byte{
	raw := make([]byte,10000)
	fmt.Println(len(raw),cap(raw),&raw[0])
	res := make([]byte,3)
	copy(res,raw[:3])
	return res
}

func main() {
	//BAD
	//data := getBad()
	//GOOD
	data := getGood()
	fmt.Println(len(data),cap(data),&data[0]) // выводит: 3 10000 <byte_addr_x>
}
	

