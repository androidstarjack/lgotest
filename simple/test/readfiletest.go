package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const FilePath = "E:\\englis_install\\go_pro\\src\\com.yuer.gio\\lgotest\\simple\\tempfile\\"

//使用io/ioutil进行读写文件
//https://blog.csdn.net/wangshubo1989/article/details/74777112/
func ioUtilsReadFile(finePath string){

	/***************************************************************************************************************/
	byteFile ,err := ioutil.ReadFile(finePath)
	fmt.Printf("byteFile: %s",string(byteFile))
	return
	//通过 ReadAll() 读取 r 中的所有数据，返回读取的数据和遇到的错误。
	file,_ := os.Open(finePath)
	file2 ,err  := ioutil.ReadAll(file)

	//通过 ioutil.ReadFile
	byteFile ,err  = ioutil.ReadFile(finePath)

	//读取文件夹
	var files  [] os.FileInfo
	files ,_ = ioutil.ReadDir(finePath)
	for i, file := range files {
		fmt.Printf("啦啦啦，遍历出第%d个文件: %s",i,file.Name())
		fmt.Println()
	}

	r := bufio.NewReader(file)
	b, err := r.Peek(1)
	fmt.Println("文件内容： ",b)


	/***************************************************************************************************************/
	// TempFile 在目录 dir 中创建一个临时文件并将其打开
	// 文件名以 prefix 为前缀+时间戳的文件夹
	// 返回创建的文件的对象和创建过程中遇到的任何错误
	// 如果 dir 为空，则在系统的临时目录中创建临时文件
	// 如果环境变量中没有设置系统临时目录，则在 /tmp 中创建临时文件
	// 调用者可以通过 f.Name() 方法获取临时文件的完整路径
	// 调用 TempFile 所创建的临时文件，应该由调用者自己移除 eg: defer os.RemoveAll(dir) // clean up
	content := []byte("temporary file's content,哈哈哈哈哈，来一起啪go语言")
	dir, err := ioutil.TempDir(FilePath, "example")
	if err != nil {
		panic(err)
	}
	fmt.Printf("创建一个缓存出第个文件: %s",dir)
	//defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		panic(err)
	}
	fmt.Printf("最终的结果: %s",tmpfn)
	/***************************************************************************************************************/



	fmt.Println(byteFile,file,file2,err,files)
}




func main1() {

}
func main(){
	//ioUtilsReadFile(FilePath)
	//writeWay()
	//ioUtilsReadFile(FilePath+"abc.yuer")
	ioUtilsReadFile(FilePath+"abc.yuer")
}

//写文件的方式
func writeWay()  {

	file, err := os.Create(FilePath+"abc.yuer2")

	if err != nil{
		panic(err)
		return
	}

	//通过ioutil的writeFile
	filePath := FilePath+"abc.yuer"
	byteData := [] byte("来来来，一起啪啪go语言") //文件不存在，自动创建
	error := ioutil.WriteFile(filePath,byteData,os.ModeAppend)
	if error != nil{
		panic(error)
	}

	//通过os进行写入
	//file ,err := os.Open(filePath)
	file ,err  = os.OpenFile(filePath ,os.O_CREATE|os.O_APPEND|os.O_RDWR,0660  )//需要提供文件路径、打开模式、文件权限
	if err != nil{
		panic(err)
	}
	defer file.Close()
	file.Write(byteData)
	file.WriteString("\n")
	n, err := file.WriteString("来来来，一起来")

	if err != nil {
		panic(err)
	}
	fmt.Printf("n : %d",n)
	defer  file.Close()

}