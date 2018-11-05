package main

import (
	"bufio"
	"com.yuer.gio/lgotest/queue/entry"
	"fmt"
	"strings"
)





func main() {
	fmt.Print("================================== \n")
	fmt.Print("================================== \n")
	node := queue.Node{}

	fmt.Printf("node: %d \n",node)
	node.Push(1)
	node.Push(2)
	node.Push(3)
	node.Push(4)
	fmt.Printf("node: %v \n",node)
	node.Push(5)
	fmt.Println(node.Pop())
	fmt.Println(node.Pop())
	fmt.Println(node.Pop())
	fmt.Println(node.Pop())
	fmt.Println(node.Pop())

	fmt.Printf("node: %v \n",node)
	fmt.Printf("node: %v \t\n",node.IsEmpety())

	scanner:=bufio.NewScanner(
		strings.NewReader("ABCDEFG\nHIJKELM"),
	)
	scanner.Split( nil /*四种方式之一，你也可以自定义, 实现SplitFunc方法*/ )
	for scanner.Scan(){
		fmt.Println(scanner.Text()) // scanner.Bytes()
	}

}
