package main

import (
	"fmt"
	"sync"
)

type automaticInt struct {
	i int
	lock sync.Mutex
}


var oridinaryInt int

var channelInt chan int


//type atomaticInt int

func(a *automaticInt) add(){
	a.lock.Lock()
	defer  a.lock.Unlock()
	a.i ++
}

func (a * automaticInt) get() int{
	a.lock.Lock()
	defer  a.lock.Unlock()
	return  a.i
}

func main() {
	//go func() {
	//	fmt.Printf("automaticInt: %d",automaticInt)
	//	for   {
	//		automaticInt.lock.Lock()
	//		automaticInt.i++
	//		automaticInt.lock.Unlock()
	//		fmt.Printf("automaticInt: %d",	automaticInt.i)
	//		if 	automaticInt.i > 1000{
	//			break
	//		}
	//	}
	//
	//}()

	var  a automaticInt

	a.add()
	go func() {
		a.add()
	}()
	//time.Sleep(time.Microsecond)
	fmt.Printf("得到的结果: automaticInt: %d",a.get())
}
















