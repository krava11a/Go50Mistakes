package main

import (
	"fmt"
	"time"
)

func main() {
	//WRONG
	//var ch chan int
	//
	//for i := 0; i < 3; i++ {
	//	go func(idx int) {
	//		ch <- (idx + 1) * 2
	//	}(i)
	//}
	//
	//fmt.Println("result : ",<-ch)
	//time.Sleep(2 * time.Second)

	//BETTER
	inch := make(chan int)
	outch := make(chan int)

	go func() {
		var in <-chan int = inch
		var out chan<- int
		var val int
		for {
			select {
			case out <- val:
				out = nil
				in = inch
			case val = <-in:
				out = outch
				in = nil
			}
		}
	}()

	go func() {
		for r := range outch {
			fmt.Println("result", r)
		}
	}()

	time.Sleep(0)
	inch <- 1
	inch <- 2
	time.Sleep(3*time.Second)
}
