package main

import (
	"fmt"
	"os"
)

const filePathDir = "E:\\englis_install\\go_pro\\src\\com.yuer.gio\\lgotest\\simple\\tempfile\\"

func main() {
	//fmt.Fscanf()
	//fmt.Fscan() //和Fscanf一样，格式化蚕食默认为空格并将数据以空格为分割符进行分割
	//fmt.Fscanln()
	//
	//fmt.Fprintf()
	//fmt.Fprint()
	//fmt.Fprintln()


	goFscanfTest()
	goFscanTest()
	goFscanlnTest()

	goFprintlnTest()
	goFprintTest()
	goFprintfTest()
}
func goFscanfTest() {
	fmt.Println("===========================fmt进行对文件操作的学习笔记====================================")
	//func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
	// Fscanf scans text read from r, storing successive space-separated
	// values into successive arguments as determined by the format. It
	// returns the number of items successfully parsed.
	// Newlines in the input must match newlines in the format.
	//Fscanf 方法 扫描r文件，解析连续的空格分离 ，参数格式一定要确定
	//返回值是连续的参数解析项
	// Fscanf 用于扫描 r 中的数据，并根据 format 指定的格式
	// 将扫描出的数据填写到参数列表 a 中
	// 当 r 中的数据被全部扫描完毕或扫描长度超出 format 指定的长度时
	// 则停止扫描（换行符会被当作空格处理）
	file ,err :=os.Open(filePathDir+"fmt.yuer")
	if err != nil{
		panic(err)
	}
	defer file.Close()
	var language string
	var grade int
	fmt.Fscanf(file,"一起来pk %s 吧，今天你考试能 %d 分吗\n",&language,&grade)
	fmt.Println(language,grade)


}

func goFscanTest(){
	fmt.Println("===========================fmt Fscan进行对文件操作的学习笔记====================================")
	//和Fscanf一样，格式化params默认为空格并将数据以空格为分割符进行分割
	// Fscan 用于扫描 r 中的数据，并将数据以空格为分割符进行分割
	// 然后填写到参数列表 a 中
	// 当 r 中的数据被全部扫描完毕或者参数列表 a 被全部填写完毕
	// 则停止扫描（换行符会被当作空格处理）
	file ,err :=os.Open(filePathDir+"fmt.yuer")
	if err != nil{
		panic(err)
	}
	defer file.Close()
	//content := [...] string{}
	var content1  string
	var content2  string
	var content3  string
	var content4  string
	var content5  string
	result ,erro := fmt.Fscan(file,&content1,&content2,&content3,&content4,&content5)
	if erro != nil {
		panic(erro)
	}
	fmt.Printf("返回解析数据字段解析的结果result :%d \n",result)
	fmt.Printf("content :%s  %s  %s  %s  %s  \n",content1,content2,content3,content4,content5)
}

func goFscanlnTest(){
	fmt.Println("===========================fmt Fscanln进行对文件操作的学习笔记====================================")
	//和Fscanf一样，格式化params默认为空格并将数据以空格为分割符进行分割
	// Fscan 用于扫描 r 中的数据，并将数据以空格为分割符进行分割
	// 然后填写到参数列表 a 中
	// 当 r 中的数据被全部扫描完毕或者参数列表 a 被全部填写完毕
	// 则停止扫描（换行符会被当作空格处理）
	//Fscanln 将一次只读一行，并将每列的数据赋值给相应的变量
	file ,err :=os.Open(filePathDir+"fmt.yuer")
	if err != nil{
		panic(err)
	}
	defer file.Close()
	//content := [...] string{}
	var content1  string
	var content2  string
	var content3  string
	var content4  string
	var content5  string
	result ,erro := fmt.Fscanln(file,&content1,&content2,&content3,&content4,&content5)

	if erro != nil {
		panic(erro)
	}
	fmt.Printf("返回解析数据字段解析的结果result :%d \n",result)
	fmt.Printf("content :%s  %s  %s  %s  %s  \n",content1,content2,content3,content4,content5)
}

func goFprintlnTest(){
	fmt.Println("===========================fmt Fprintln进行对文件操作的学习笔记====================================")
	//指定数据，向文件中写入
	//Fprintln 将一次只写入一行，并将每列的数据赋值给相应的变量
	file ,err :=os.OpenFile(filePathDir+"fmt.yuer",os.O_APPEND|os.O_RDWR|os.O_SYNC,0660)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	//content := [...] string{}
	var content1  = "来"
	var content2   = "一"
	var content3  = "起"
	var content4  = "啪啪啪"
	var content5  = "Go"
	result ,erro := fmt.Fprintln(file,content1,content2,content3,content4,content5)

	if erro != nil {
		panic(erro)
	}
	fmt.Printf("返回解析数据字段解析的结果result :%d \n",result)
	fmt.Printf("content :%s  %s  %s  %s  %s  \n",content1,content2,content3,content4,content5)
}
func goFprintTest(){
	fmt.Println("===========================fmt Fprint进行对文件操作的学习笔记====================================")
	//指定数据，向文件中写入 Fprint 对行数没有限制，并将每列的数据赋值给相应的变量
	file ,err :=os.OpenFile(filePathDir+"fmt.yuer",os.O_APPEND|os.O_RDWR|os.O_SYNC,0660)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	//content := [...] string{}
	var content1  = "来"
	var content2   = "一"
	var content3  = "起"
	var content4  = "啪啪啪"
	var content5  = "Go"
	var content6  = "Goasdffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffasdfasdfasdfasdfasdf\nasdfasdfaosdfasdfapsdfasdfsds"
	result ,erro := fmt.Fprint(file,content1,content2,content3,content4,content5,content6)

	if erro != nil {
		panic(erro)
	}
	fmt.Printf("返回解析数据字段解析的结果result :%d \n",result)
	fmt.Printf("content :%s  %s  %s  %s  %s  \n",content1,content2,content3,content4,content5)
}
func goFprintfTest(){
	fmt.Println("===========================fmt Fprintf进行对文件操作的学习笔记====================================")
	//指定数据，向文件中写入
	// Fprintf 将参数列表 a 填写到格式字符串 format 的占位符中
	// 并将填写后的结果写入 w 中，返回写入的字节数
	file ,err :=os.OpenFile(filePathDir+"fmt.yuer",os.O_APPEND|os.O_RDWR|os.O_SYNC,0660)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	//content := [...] string{}
	var content1  = "来"
	var content2   = "一"
	var content3  = "起"
	var content4  = "啪啪啪"
	var content5  = "Go"
	result ,erro := fmt.Fprintf(file,"fmt Fprint进行对文件操作的学习笔记 %s  %s  %s  %s  %s ",content1,content2,content3,content4,content5)

	if erro != nil {
		panic(erro)
	}
	fmt.Printf("返回解析数据字段解析的结果result :%d \n",result)
	//asdfas
}