package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const filePath = "E:\\englis_install\\go_pro\\src\\com.yuer.gio\\lgotest\\simple\\tempfile\\"

func readFileTest(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//逐行读取有的时候真的很方便，性能可能慢一些，但是仅占用极少的内存空间。
	buf := make([]byte, 100)
	res, ero := file.Read(buf)
	if ero != nil {
		panic(ero)
	}

	fileData := string(buf)
	fmt.Printf("结果：%d  \n 数据： \n  %s", res, fileData)
	fmt.Println("=============================")

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	_, err = reader.Read(buffer)
	fmt.Printf("最终文件结果>>>：%s", string(buffer))
	fmt.Println("最终文件结果结束")
	//fmt.Printf("-------------> %-20s %-2v %v\n", buffer[:n], n, err)
	// ABCDEFGHIJKLMNOPQRST 20 <nil>

	//小文件推荐一次性读取，这样程序更简单，而且速度最快。
	fileData1, err := ioutil.ReadAll(file)
	check(err)
	fileData = string(fileData1)
	fmt.Printf("结果：  \n 数据： \n  %s", fileData)
	fmt.Println("=============================")

	//加简单的方法
	fileData1, err = ioutil.ReadFile(filePath)
	fileData = string(fileData1)
	fmt.Printf("结果：  \n 数据： \n  %s", fileData)
	check(err)
	fmt.Println("=============================")

}

func readDir() {

	/***************************************************************************************************************/
	byteFile, err := ioutil.ReadFile(filePath)
	fmt.Printf("byteFile: %s", string(byteFile))
	return
	//通过 ReadAll() 读取 r 中的所有数据，返回读取的数据和遇到的错误。
	file, _ := os.Open(filePath)
	file2, err := ioutil.ReadAll(file)

	//通过 ioutil.ReadFile
	byteFile, err = ioutil.ReadFile(filePath)

	//读取文件夹
	var files []os.FileInfo
	files, _ = ioutil.ReadDir(filePath)
	for i, file := range files {
		fmt.Printf("啦啦啦，遍历出第%d个文件: %s", i, file.Name())
		fmt.Println()
	}

	r := bufio.NewReader(file)
	b, err := r.Peek(1)
	fmt.Println("文件内容： ", b)

	/***************************************************************************************************************/
	// TempFile 在目录 dir 中创建一个临时文件并将其打开
	// 文件名以 prefix 为前缀+时间戳的文件夹
	// 返回创建的文件的对象和创建过程中遇到的任何错误
	// 如果 dir 为空，则在系统的临时目录中创建临时文件
	// 如果环境变量中没有设置系统临时目录，则在 /tmp 中创建临时文件
	// 调用者可以通过 f.Name() 方法获取临时文件的完整路径
	// 调用 TempFile 所创建的临时文件，应该由调用者自己移除 eg: defer os.RemoveAll(dir) // clean up
	content := []byte("temporary file's content,哈哈哈哈哈，来一起啪go语言")
	dir, err := ioutil.TempDir(filePath, "example")
	if err != nil {
		panic(err)
	}
	fmt.Printf("创建一个缓存出第个文件: %s", dir)
	//defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		panic(err)
	}
	fmt.Printf("最终的结果: %s", tmpfn)
	/***************************************************************************************************************/

	fmt.Println(byteFile, file, file2, err, files)
}

func writeFileTest() {

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	readFileTest(filePath + "fmt.yuer")
}
