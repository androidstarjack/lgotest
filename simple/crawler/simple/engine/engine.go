package engine

import (
	"com.yuer.gio/lgotest/simple/crawler/simple/fecther"
	"fmt"
	"log"
)

//import "com.yuer.gio/lgotest/simple/crawler/engine"

//engine引擎循坏执行的方法
var index = 0
func Run(seeds ...Request){

	log.Printf("传过来的参数:1 ,  %v 打印参数结束啦~~~",seeds)

	var requests [] Request

	for _ , val := range seeds{
		requests = append(requests, val)
	}

	log.Println("传过来的参数:seeds的长度",len(requests) ,"打印参数结束啦~~~\n")
	  var index = 0
	for len(requests)>0 {

		r := requests[0]
		requests = requests[1:]

		parseResult,erro := work(r)
		if erro != nil{
			continue
		}
		//fmt.Println("打印parseResult信息 :",parseResult.Items,parseResult.Requests)
		fmt.Printf("Requests的长度： %d :", len(requests))
		requests = append(requests,parseResult.Requests...)
		//fmt.Println("打印prequests信息 :",requests)
		index = index +1
		//log.Panicf("这是第 %d 次 执行解析。，解析的结果为： %v",index,requests)
		fmt.Printf("这是第 %d 次 执行解析-----解析的结果为： %d \n",index, len(requests))
		for _, items := range  parseResult.Items{
			index ++
			log.Printf("got item  %d data : %s",index,items)
		}

		//fmt.Printf("集合获取的长度： %d\n",len(requests))



	}
}



func work(r Request,) (ParseResult, error){
	body ,erro := fecther.ParseUrlByUtf8Tansfer(r.Url)
	//log.Printf("Fecthing: -------- - - - - - - -  - ->%s",r.Url)
	//fmt.Println("")
	//fmt.Println("啊啊啊啊啊啊啊啊啊啊啊",string(body),"结束啦结束啦")

	if erro != nil{
		log.Printf("ParseUrlByUtf8Tansfer erro " +
			"feacth url %s ,%s",r.Url,erro)
		return ParseResult{},erro
	}
	return  r.ParseFunc(body) ,nil
}






func Run1(seeds Request){

	log.Println("传过来的参数:1",seeds ,"打印参数结束啦~~~")

	var requests [] Request

	//for _ , val := range seeds{
	//	requests = append(requests, val)
	//}
	requests = append(requests, seeds)

	log.Println("传过来的参数:seeds的长度",len(requests) ,"打印参数结束啦~~~")

	//for len(requests)>0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fecthing: %s",r.Url)
		data ,erro := fecther.ParseUrlByUtf8Tansfer(r.Url)
		if erro != nil{
			log.Printf("ParseUrlByUtf8Tansfer erro " +
				"feacth url %s ,%s",r.Url,erro)
			//continue
		}
	log.Printf("Fecthing: %s------》： %s",r.Url,string(data))
		parseResult := r.ParseFunc(data)
		requests = append(requests,parseResult.Requests...)
		//requests =
		for _, items := range  parseResult.Items{
			log.Printf("got item data : %s",items)
			//log.Print(items)
		}
		for _, items := range  parseResult.Requests{
		log.Printf("got item data : %s",items.ParseFunc)
		//log.Print(items)
		}
	//}
	//最后再进行解析

}