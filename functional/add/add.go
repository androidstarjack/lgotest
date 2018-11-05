package main

import (
	"fmt"

)

func add() func(value int )int {
	 sum := 0 
	 return func(value int) int {
		 sum += value
		 return   sum
	 }
}


func main() {
	a :=add()
	//ss := 0
	//fmt.Println(ss(1))
	for i:=0 ; i< 10; i ++{
		fmt.Printf("0+....+  %d = %d\n",i ,a(i))
	}
}
