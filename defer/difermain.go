package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const FILEPATH ="abcTest1.txt"

func deferTest(filePath string){
	file, error := os.OpenFile(filePath, os.O_EXCL|os.O_CREATE, 0666)
	if error != nil{
		error = errors.New("is not a pathError")
		if pathError , ok :=error.(*os.PathError); !ok{

			fmt.Printf("最终显示： %t",error)
			panic(error)
		}else{
			fmt.Println("ok为true的时候 ： ",pathError.Op,pathError.Path,pathError.Err)
			fmt.Printf("最终显示： %s  %s %s",pathError.Op,pathError.Path,pathError.Err)
		}
		fmt.Printf("最终显示： %s  %s %s",error)
		return
	}
	defer  file.Close()
	writer := bufio.NewWriter(file)
	//writer.WriteString("你大爷的\nDNS查询也消耗响应时间，如果我们的网页内容来自各个不同的domain (比如嵌入了开放广告，引用了外部图片或脚本)，那么客户端首次解析这些domain也需要消耗一定的时间。DNS查询结果缓存在本地系统和浏览器中一段时间，所以DNS查询一般是对首次访问响应速度有所影响。下面是我清空本地dns后访问博客园主页dns的查询请求。")
	defer writer.Flush()

	f := Feibinaci()

	for i:=0; i < 20; i++{
		fmt.Fprintln(writer,f())
	}

}
type genAdd func() int
func  Feibinaci() genAdd{
	a,b := 0,1
	return func() int {
		a,b  = b ,a+b
		return  a
	}
}


func main() {
	deferTest(FILEPATH)
}
