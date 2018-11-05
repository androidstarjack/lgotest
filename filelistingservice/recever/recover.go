package main

import "fmt"

func tryrecover(){
	defer func(){
		r := recover()

		if err ,ok  := r.(error); ok{
			fmt.Println("error occurred: " ,err.Error())
		}else{
			if err != nil{
				panic(err)
			}else{
				fmt.Print(err.Error())
				panic("我不知道该如何去处理错误啦 ")
			}
		}
	}()

	//b := 0
	//c := 5 / b
	//fmt.Println(c)
	panic(123)
	fmt.Print("我在最后，看来我开始执行啦~~~~")
}
func main() {
	tryrecover()
}
