package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createChanInt () chan int{
	c := make(chan int)
	go func() {
		i := 0
		for{
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Microsecond)
			c <- i
			i++
		}
	}()
	return c
}


func funcTest3(){
	var c1 ,c2 =createChanInt(),createChanInt()
	for   {
		select {
		case n:= <-c1:
			fmt.Println("receiver: data：  from c1",n)
		case n:=<-c2:
			fmt.Println("receiver: data：  from c2",n)

		}
	}
}


func main(){
	funcTest3()
}

func main2() {
	ch := make(chan int)
	go func() {
		var sum int = 100
		fmt.Println("第一次发送龟派气功")
		ch <- sum
		fmt.Println("第二次即将发送")
		fmt.Println("第二次发送气功")
		ch <- sum
	}()
	time.Sleep(1e9)
	fmt.Println("第一次中招了")
	fmt.Println(<-ch)
	fmt.Println("第二次中招了")
	fmt.Println(<-ch)
}
func main1() {
	ch := make(chan int, 5)
	go func() {
		var sum int = 6
		fmt.Println("第一次发送龟派气功")
		ch <- sum
		fmt.Println("第二次即将发送")
		time.Sleep(1e9)
		fmt.Println("第二次发送气功")
		ch <- sum
		fmt.Println("第三次即将发送")
		time.Sleep(1e9)
		fmt.Println("第三次发送气功")
		ch <- sum
		time.Sleep(1e9)
		fmt.Println("第四次发送气功")
		ch <- sum
	}()
	time.Sleep(6e9)
	fmt.Println("第一次中招了")
	fmt.Println(<-ch)
	fmt.Println("第二次中招了")
	fmt.Println(<-ch)
	fmt.Println("第三 次中招了")
	fmt.Println(<-ch)
	fmt.Println("第四次中招了")
	fmt.Println(<-ch)
}

