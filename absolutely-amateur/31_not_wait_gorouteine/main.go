package main

import (
	"fmt"
	"sync"
	"time"
)

func testGoRoutine(number int) {
	fmt.Println("Start goroutine ", number)
	time.Sleep(1 * time.Second)
	fmt.Println("stop goroutine", number)
}

func main() {

	//WRONG WAY
	//for i := 0;i<3;i++ {
	//	go testGoRoutine(i)
	//}
	//time.Sleep(1*time.Second)
	//fmt.Println("main exit")
	//

	//With WG and Channel
	//var wg sync.WaitGroup
	//done := make(chan struct{})
	//workerCount := 1002
	//wg.Add(workerCount)
	//for i := 0; i < workerCount; i++ {
	//
	//	go testGoRoutineWithChannel(i, done, &wg)
	//}
	//
	//close(done)
	//wg.Wait()
	//fmt.Println("All done!!")

	//With WG,Channel and Select
	var wg sync.WaitGroup
	done := make(chan struct{})
	wq := make(chan interface{})
	workerCount := 5

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go testGoRoutineWithChannelAndSelect(i,wq, done, &wg)
	}

	for i := 0; i < workerCount; i++ {
		wq<-i
	}

	close(done)
	wg.Wait()
	fmt.Println("All done!!")


}

func testGoRoutineWithChannel(workerId int, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerId)
	defer wg.Done()
	<-done
	fmt.Printf("[%v] is done\n", workerId)

}

func testGoRoutineWithChannelAndSelect(workerId int, wq <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerId)
	defer wg.Done()
	for {
		select {
		case m := <-wq:
			fmt.Printf("[%v] m => %v\n", workerId, m)
		case <-done:
			fmt.Printf("[%v] is done\n", workerId)
			return
		}
	}


}
