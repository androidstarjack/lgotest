package main

import (
	"com.yuer.gio/lgotest/simple/crawler/engine"
	"com.yuer.gio/lgotest/simple/crawler/zhenai/parse"
)

const URL  =  "http://www.zhenai.com/zhenghun"



func main() {
	engine.Run(engine.Request{
		Url:URL,
		ParseFunc:parse.ParseCityListdata,
	})
	//resp ,erro :=http.Get(URL)
	//if erro != nil{
	//		panic(erro)
	//}
	//defer  resp.Body.Close()
	//
	//if resp.StatusCode != http.StatusOK{
	//	fmt.Printf("request failt,statuscode is %d",resp.StatusCode)
	//	return
	//}
	////转判断charset编码格式
	//
	//// 使用正则表达式进行过滤我们想要的字段
	//encod ,nReader := tansformGetUtf8(resp.Body)
	//reader := transform.NewReader(nReader,encod.NewEncoder())
	//data,ero :=ioutil.ReadAll(reader)
	//if ero != nil{
	//	panic(ero)
	//}
	////fmt.Println(string(data))
	//printCityListdata(data)
}
