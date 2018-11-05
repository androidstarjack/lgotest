package main

import (
	"fmt"
	"os"
)
const flePathDir = "E:\\englis_install\\go_pro\\src\\com.yuer.gio\\lgotest\\simple\\tempfile\\"
//创建文件
func main() {
	createFile()
}

func createFile(){
	filePath :=  flePathDir+"crateFile.yuer"
	fmt.Println(filePath)
	file, erro := os.Create(filePath)
	if erro != nil{
		panic(erro)
	}
	defer  file.Close()
	file.WriteString("创建一个文件~~~")
}