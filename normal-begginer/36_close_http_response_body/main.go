package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// MAY BE
	//resp, err := http.Get("https://api.ipify.org/?format=json")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer resp.Body.Close() // допустимо, в большинстве случаев :-)
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))


	//BETTER
	resp, err := http.Get("https://api.ipify.org?format=json")
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
