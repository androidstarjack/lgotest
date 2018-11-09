package fecther

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)
const URL  =  "http://www.zhenai.com/zhenghun"

func ParseUrlByUtf8Tansfer(url string) ([] byte,error ){
	//if strings.Contains(url,"http://www.zhenai.com/"){
	//	return ParseUrlByUtf8TansferGet(url)
	//} else if strings.Contains(url,"http://album.zhenai.com/u"){
	//	return ParseUrlByUtf8TansferPost(url)
	//}
	return ParseUrlByUtf8TansferGet(url)
}
func ParseUrlByUtf8TansferGet(url string) ([] byte,error ){
	resp ,erro :=http.Get(url)
	if erro != nil{
	 return nil,fmt.Errorf("request url error :%s",erro)
	}
	defer  resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Printf("request failt,statuscode is %d",resp.StatusCode)
		return nil,fmt.Errorf("request failt,statuscode is %d",resp.StatusCode)
	}
	//转判断charset编码格式

	// 使用正则表达式进行过滤我们想要的字段
	encod ,nReader := tansformGetUtf8(resp.Body)
	reader := transform.NewReader(nReader,encod.NewEncoder())
	return ioutil.ReadAll(reader)

	//data ,erro := ioutil.ReadAll(reader)
	//if erro != nil{
	//	panic(erro)
	//}
	//fmt.Println("")
	//fmt.Println("啊啊啊啊啊啊啊啊啊啊啊",string(data))
	//printCityListdata(data)
	//return ioutil.ReadAll(reader)
}
func ParseUrlByUtf8TansferPost(url string) ([] byte,error ){ //

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}


	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.26 Safari/537.36 Core/1.63.5702.400 QQBrowser/10.2.1893.400")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil{
	 return nil,fmt.Errorf("request url error :%s",err)
	}
	defer  resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Printf("request failt,statuscode is %d",resp.StatusCode)
		return nil,fmt.Errorf("request failt,statuscode is %d",resp.StatusCode)
	}
	//转判断charset编码格式

	// 使用正则表达式进行过滤我们想要的字段
	encod ,nReader := tansformGetUtf8(resp.Body)
	reader := transform.NewReader(nReader,encod.NewEncoder())
	data ,erp := ioutil.ReadAll(reader)
	if erp != nil{
		panic(erp)
	}
	fmt.Println("结果： -----》",string(data))
	return data,erp

	//printCityListdata(data)
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