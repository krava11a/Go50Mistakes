package main

import (
	"fmt"
	"runtime"
)

func main() {

	//maximum - 256
	fmt.Println(runtime.GOMAXPROCS(-1)) // выводит: X (1 on play.golang.org)
	fmt.Println(runtime.NumCPU())       // выводит: X (1 on play.golang.org)
	runtime.GOMAXPROCS(20)
	fmt.Println(runtime.GOMAXPROCS(-1)) // выводит: 20
	runtime.GOMAXPROCS(300)
	fmt.Println(runtime.GOMAXPROCS(-1)) // выводит: 256
}
