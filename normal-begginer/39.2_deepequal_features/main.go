package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {

	var b1 []byte = nil
	b2 := []byte{}
	fmt.Println("b1==b2", reflect.DeepEqual(b1, b2))

	var str string = "one"
	var in interface{} = "one"
	fmt.Println("str == in", str == in, reflect.DeepEqual(str, in))

	v1 := []string{"one", "two"}
	v2 := []interface{}{"one", "two"}
	fmt.Println("v1 == v2", reflect.DeepEqual(v1, v2)) // prints true, and this not OK

	data := map[string]interface{}{
		"code":  200,
		"value": []string{"one", "two"},
	}
	encoded, _ := json.Marshal(data)
	var decoded map[string]interface{}
	json.Unmarshal(encoded, &decoded)
	fmt.Println("data == decoded", reflect.DeepEqual(data, decoded))

}
