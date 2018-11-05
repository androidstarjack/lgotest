package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	deleteFile()
}
func deleteFile() {
	fileDIrPath := "E:\\englis_install\\go_pro\\src\\com.yuer.gio\\lgotest\\simple\\tempfile\\"+"crateFile.yuer"
	file,err := os.Open(fileDIrPath)
	 if  err != nil{
			panic(err)
	 }

	fmt.Println("检测到该文件")
	data ,erro :=ioutil.ReadFile(fileDIrPath)
	if erro != nil{
		panic(erro)
	}
	fmt.Printf("文件：%s",string(data))
	file.Close()
	//先关掉file，不然会报错的，提示进程一直在使用，无法访问

	//err  = os.Remove(fileDIrPath)
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Println("删除文件成功")



	// 如果文件不存在，则返回错误
	fileInfo, err := os.Stat(fileDIrPath)
	if err != nil{
		panic(err)
	}
	fmt.Println(fileInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}
