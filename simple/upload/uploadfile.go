package main

import (
	"bufio"
	"com.yuer.gio/lgotest/simple/consts"
	"crypto/md5"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

//func uploadFile()
func main() {

	http.HandleFunc("/uploadfile",uploadFile)
	http.ListenAndServe(":9090",nil)
}



func uploadFile(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("mother: %s",request.Method)

	if request.Method == "GET"{
		CurrentTimer := time.Now().Unix()
		h := md5.New()
		io.WriteString(writer,strconv.FormatInt(CurrentTimer,10))
		token := fmt.Sprintf("%x",h.Sum(nil))
		t , _ :=template.ParseFiles("upload gtpl")
		t.Execute(writer,token)
	}else{
		request.ParseMultipartForm(32 << 20)


 	 // 使用正则表达式进行过滤我们想要的字段
		//encod ,nReader := tansformGetUtf8(request.Body)
		//reader := transform.NewReader(nReader,encod.NewEncoder())


		file,handler,erro := request.FormFile("aaaaa")
		if erro != nil{
			fmt.Println(erro)
		}
		defer file.Close()
		fmt.Printf("打印结果: %s  %v \n",request.Method,handler.Header)
		fmt.Fprintf(writer,"%v",handler.Header)
		f,e := os.OpenFile(consts.FilePath+handler.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if e!= nil{
			fmt.Println("文件打开出错啦",e)
			return
		}
		defer f.Close()
		_, crrl := io.Copy(f,file)
		if crrl != nil{
			println("copy错误了   ",crrl)
		}
	}
}

func tansformGetUtf8(r io.Reader)(encoding.Encoding, io.Reader) {
	//data ,erro := (bufio.NewReader(r)).Peek(1024)
	nReader := bufio.NewReader(r)
	//peek是对Reader来说的，所以应该操作的是Reader对象，而不是f
	data,erro := nReader.Peek(1024)
	if erro != nil{
		return  unicode.UTF8 ,nReader
	}
	encod,_ ,_ :=charset.DetermineEncoding(data,"")
	return  encod ,nReader
}

// 处理/upload 逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		  var  maxMemory  int64 = 32 << 20
		r.ParseMultipartForm(maxMemory)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}