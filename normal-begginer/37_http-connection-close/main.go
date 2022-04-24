package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//Чтобы не закончились сокеты или файловые дескрипторы, стоит закрывать соединение после запроса
	req, err := http.NewRequest("GET", "http://golang.org", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Close = true
	//or do this:
	//req.Header.Add("Connection","close")

	resp, err := http.DefaultClient.Do(req)
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

	fmt.Println(len(string(body)))

	//Также возможно глобально отключить повторное использование HTTP-соединения
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}
	resp, err = client.Get("http://golang.org")
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.StatusCode)
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(string(body)))

}
