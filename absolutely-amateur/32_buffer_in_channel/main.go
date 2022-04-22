package main

import (
	"fmt"
)

func main() {

	//runtime.GOMAXPROCS(1)

	ch := make(chan string)
	go func() {
		for m := range ch{
			fmt.Println("processed:",m)
		}
	}()

	ch<-"cmd.1"
	ch<-"cmd.2"  // do not work
}
