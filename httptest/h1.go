package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)






func main(){
	maminGetRequest()
}

func maminGetRequest(){
	request ,erro := http.NewRequest(http.MethodGet,"https://m.imooc.com/",nil)
	//request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect： ",req)
			return  nil
		},
	}
	resp ,mresError := client.Do(request)
	defer resp.Body.Close()
	if erro != nil {
		panic(erro)
	}
	if mresError != nil{
		panic(mresError)
	}
	resByte ,resErro := httputil.DumpResponse(resp,true)
	if resErro != nil{
		panic(resErro)
	}

	fmt.Printf("抓取到的结果： %s",resByte)
}

func maminGetRequest2(){
	resp ,erro := http.Get("https://www.baidu.com")
	if erro != nil {
		panic(erro)
	}
	defer resp.Body.Close()

	resByte ,resErro := httputil.DumpResponse(resp,true)
	if resErro != nil{
		panic(resErro)
	}

	fmt.Printf("抓取到的结果： %s",resByte)
}
func helloFunc(responseWrite http.ResponseWriter,requset * http.Request ) {
	io.WriteString(responseWrite,"你好啊，卡卡罗特")
}
func mainRequest() {
	http.HandleFunc("/list",helloFunc)
	http.ListenAndServe(":8080",nil)
}
