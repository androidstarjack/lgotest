package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type genAdd func() int

func (g genAdd) Read(p []byte) (n int, err error) {
	//panic("implement me")
	next := g()
	if next > 10000{
		return  0 ,io.EOF
	}
	s := fmt.Sprintf("%d \n",next)
	return strings.NewReader(s).Read(p)
}

func  feibinaci() genAdd{
	a,b := 0,1
	return func() int {
		a,b  = b ,a+b
		return  a
	}
}
func  Feibinaci() genAdd{
	a,b := 0,1
	return func() int {
		a,b  = b ,a+b
		return  a
	}
}

func printFileCount(read io.Reader){
	scanner := bufio.NewScanner(read)
	for scanner.Scan(){
		fmt.Printf(scanner.Text())
	}
}


func getReadFlib() string {

	return ""
}




func main() {
	f := feibinaci()
	printFileCount(f)
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())
}
