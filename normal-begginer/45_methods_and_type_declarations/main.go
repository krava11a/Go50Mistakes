package main

import (
	"fmt"
	"sync"
)

type MyMutex sync.Mutex

type MyMutexikBetter struct {
	mutexik sync.Mutex
}

func New() *MyMutexikBetter {
	fmt.Println("Like a constructor")
	return &MyMutexikBetter{}
}

func (mm *MyMutexikBetter) Lock() {
	fmt.Println("LOCK")
	mm.mutexik.Lock()
}

func (mm *MyMutexikBetter) Unlock() {
	fmt.Println("UNLOCK")
	mm.mutexik.Unlock()
}

type MyBestMutex struct {
	sync.Mutex
}

type myLocker sync.Locker

func main() {
	//BAD
	//var mtx MyMutex
	//mtx.Lock()   //mistake
	//mtx.Unlock() //mistake


	//BETTER
	//mtx:=new(MyMutexikBetter)
	//fmt.Println(&mtx)
	//mtx = New()
	//fmt.Println(&mtx)
	//mtx.Lock()
	//mtx.Unlock()


	//RIGHT
	mtx := new(MyBestMutex)
	mtx.Lock()
	mtx.Unlock()


	//But interface working
	locka := new(sync.Mutex)
	locka.Lock()
	locka.Unlock()


}
