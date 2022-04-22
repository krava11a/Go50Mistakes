package main

import (
	"fmt"
	"time"
)

func main() {

	//WRONG
	//ch := make(chan int)
	//for i := 0; i < 3; i++ {
	//	go func(idx int) {
	//		ch <- (idx + 1) * 2
	//	}(i)
	//}
	////get the first result
	//fmt.Println(<-ch)
	//close(ch)
	//time.Sleep(2 * time.Second)

	//RIGHT
	ch := make(chan int)
	done := make(chan struct{})
	for i := 0; i < 6; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println(idx, "sent result")

			case <-done:
				fmt.Println(idx, "exiting!!   ",len(done))
			}
		}(i)
	}

	fmt.Println("result: ", <-ch)
	close(done)
	time.Sleep(3 * time.Second)
	fmt.Println("All done!")
}
