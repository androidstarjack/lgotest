package main

import (
	"bufio"
	"com.yuer.gio/lgotest/simple/crawler/simple/fecther"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

const URL  =  "http://www.zhenai.com/zhenghun"

func main() {
	//testCityname()
	test2()
}
func test2() {
	_ ,erro :=fecther.ParseUrlByUtf8TansferPost("http://album.zhenai.com/u/1572218980")

	if erro != nil{
		 panic(erro)
	}
	//fmt.Print(string(resp))
}
const cityRegexp1 =`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
func testCityname() {
	city := `<a href="http://album.zhenai.com/u/1572218980" target="_blank">安然</a>`
	regxp := regexp.MustCompile(cityRegexp1)
	result   := regxp.FindAllString(city,-1)
	fmt.Printf("结果： %s",result)
}


func mainq() {
	resp ,erro :=http.Get(URL)
	if erro != nil{
			panic(erro)
	}
	defer  resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Printf("request failt,statuscode is %d",resp.StatusCode)
		return
	}
	//转判断charset编码格式

	// 使用正则表达式进行过滤我们想要的字段
	encod ,nReader := tansformGetUtf8(resp.Body)
	reader := transform.NewReader(nReader,encod.NewEncoder())
	data,ero :=ioutil.ReadAll(reader)
	if ero != nil{
		panic(ero)
	}
	fmt.Println(string(data))
	//printCityListdata(data)
}


func tansformGetUtf8(r io.Reader)(encoding.Encoding, io.Reader) {
	//data ,erro := (bufio.NewReader(r)).Peek(1024)
	nReader := bufio.NewReader(r)
	//peek是对Reader来说的，所以应该操作的是Reader对象，而不是f
	data,erro := nReader.Peek(1024)
	if erro != nil{
		panic(erro)
	}
	encod,_ ,_ :=charset.DetermineEncoding(data,"")
	return  encod ,nReader
}

func printCityListdata(data [] byte){
	regexpp := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z][A-z]+)"[^>]*>([^<]+)</a>`)
	result := regexpp.FindAllSubmatch(data,-1)

	for _ ,val := range result{
	//	fmt.Printf("解析出第 %d 条数据 : %s\n",index,val )
	//	for _, match := range val{
			fmt.Printf("城市：%s URL: %s \n",val[2],val[1] )
		//}
	}
	fmt.Printf("长度：%d",len(result))
}