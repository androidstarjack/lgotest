package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const Name  = "yuer"//类似于Java中的final 不可变变量
var MyName = "yuer"//全局静态变量

func s(){
	//如何定义变量
	var name string
	var name1   = "yuer"
	//直接赋值
	name2:="yuer"
	substr:= "y"
	isContact  := strings.Contains(name2,substr)
	if isContact{
		fmt.Print("name2中包含y字母\n")
	}
	//支持群赋值
	var q1,q2,q3,q4 = "q1","q2","q3","q4"
	fmt.Println(name,name1,name2)

	fmt.Println(q1,q2,q3,q4)
}
func Arrays(){

	var array  [] int //默认为空数组

	//第一种
	//var <数组名称> [<数组长度>]<数组元素>
	var arr1 [2]int
	arr1[0]=1
	arr1[1]=2

	//第二种
	//var <数组名称> = [<数组长度>]<数组元素>{元素1,元素2,...}
	var arr2 = [2]int{1,2}
	//或者
	arr3 := [2]int{1,2}

	//第三种
	//var <数组名称> [<数组长度>]<数组元素> = [...]<元素类型>{元素1,元素2,...}
	var arr4 = [...]int{1,2}
	//或者
	arr5:= [...]int{1,2}

	//第四种
	//var <数组名称> [<数组长度>]<数组元素> = [...]<元素类型>{索引1:元素1,索引2:元素2,...}
	var arr6 = [...]int{1:1,0:2}
	//或者
	arr7 := [...]int{1:1,0:2}

	fmt.Println(array,arr1,arr2,arr3,arr4,arr5,arr6,arr7)

}

func sliceTest(){
	var ss []string;
	fmt.Printf("[ local print ]\t:\t length:%v\taddr:%p\tisnil:%v\n",len(ss),ss, ss==nil)
	print("func print",ss)
	//切片尾部追加元素append elemnt
	for i:=0;i<10;i++{
		ss=append(ss,fmt.Sprintf("s%d",i));
	}
	fmt.Printf("[ local print ]\t:\tlength:%v\taddr:%p\tisnil:%v\n",len(ss),ss, ss==nil)
	print("after append",ss)
	//删除切片元素remove element at index
	index:=5;
	ss=append(ss[:index],ss[index+1:]...)
	print("after delete",ss)
	//在切片中间插入元素insert element at index;
	//注意：保存后部剩余元素，必须新建一个临时切片
	rear:=append([]string{},ss[index:]...)
	ss=append(ss[0:index],"inserted")
	ss=append(ss,rear...)
	print("after insert",ss)

}


func formapTest() {
	var arrayi= [...] int{1, 2, 3, 4, 5, 6, 7, 78, 9, 10}
	for index, c := range arrayi {
		fmt.Printf("array[%d] = %d", index, c)
	}

	str := "go语言的学习和啪啪"
	for i, ch := range str {
		fmt.Println(i, ch) //ch的类型为rune unicode编码
	}
	//输出为：28907  （Unicode编码时，两个字节代表一个字符）

	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch)
		//输出为utf-8编码，一个汉字字符占三个字节
	}

	array := []rune(str)
	n = len(array)
	for i := 0; i < n; i++ {
		ch := array[i]     // 依据下标取字符串中的字符，类型为byte
		 fmt.Println(i, ch) //unicode 编码转十进制输出
		//golang中字符类型的实际数据类型为uint32,以for循环遍历的方式输出结果都是Unicode编码的
	}
	//var str string= "yyh，hello，卡卡论坛，好厉害哦"
	//fmt.Print(str)
	fmt.Print("\n================================\n")
	for i , ch :=  range str{
		//fmt.Printf("(%d, %c)",i,ch)
		fmt.Printf("(%d, %x)",i,ch)
	}

	fmt.Print(utf8.RuneCountInString(str))
	fmt.Print("================================\n")
	bytes := [] byte(str)

	//for len(bytes) > 0 {
		r, size := utf8.DecodeRune(bytes)
		fmt.Printf("%c  %d",r,size)
	//}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	for i,c := range bytes{

		r,_ :=utf8.DecodeRune(bytes)
		fmt.Printf("%d  %c %x \n",i,r,c)
	}
	for i ,ch := range []rune(str){
		fmt.Printf("%d: %c  ",i,ch)
	}

	str2 := "123 我按时 的发ad fg票 是否 adfg 发 发生a f发 的sj df"
	sps   := strings.Split(str2," ")
	sps = strings.Fields(str2)
	var isContact  =  strings.Contains(str2,"你们")
	fmt.Println(sps)
	fmt.Println(isContact)
}


//	Go里面有一个关键字iota，这个关键字用来声明enum的时候采用，它默认开始值是0，每调用一次加1：
const(
	x = iota  // x == 0
	y = iota  // y == 1
	z = iota  // z == 2
	w  // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用" )
)

func enumTest(){
	const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0
	xtype := getXandY()
	switch xtype {
	case x:
		//go语言不需要break，默认添加的有
		break
	case y:
		break
	case z:
		break
	case w:
		break
	}
}
func getXandY() int  {
	return 0
}


func mapTest(){
	var m map[string]string // 声明一个hashmap，还不能直接使用，必须使用make来初始化
	m = make(map[string]string) // 初始化一个map
	m = make(map[string]string, 3) // 初始化一个map并附带一个可选的初始bucket（非准确值，只是有提示意义）
	m2 := map[string]string{} // 声明并初始化
	m3 := make(map[string]string) // 使用make来初始化

	map1 := make(map[string]string, 5)
	map2 := make(map[string]string)
	map3 := map[string]string{}
	map4 := map[string]string{"a": "1", "b": "2", "c": "3"}


	fmt.Println(map1, map2, map3, map4,m,m2,m3)

	//map的增删改查操作
	map4["d"] = "4"
	map4["e"] = "5"
	map4["f"] = "6"

	// 如果访问一个不存在的key，返回类型默认值
	fmt.Println("结果:",map4["b"]) // 输出0

	// 测试key是否存在
	k,ok := map4["f"]
	if ok{
		delete(map4,k)
	}

}
func main() {
	//s()
	//Arrays()
	//mapTest()
	formapTest()
}


//golang中的接口学习
func interfacetest(){

}

