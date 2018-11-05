package main

import (
	"com.yuer.gio/lgotest/retriever/deal"
	"com.yuer.gio/lgotest/retriever/mock"
	"fmt"
	"time"
)

const  URL  =  "https://www.baidu.com/"

type Retriever interface  {
	Get(url string) string
}

type Poster interface {
	post(url string,
		form map[string]string) string
}

func Post(poster Poster){
	poster.post("https://www.baidu.com/",
		map[string]string{
			"name":"张三",
			"age":"15",
			"sex":"男",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
	//papapa()
}

func session(retrieverPoster RetrieverPoster) string{
	retrieverPoster.post(URL, map[string]string{
		"name":"yuer",
		"age":"27",
		"sex":"y",
	})
	return retrieverPoster.Get(URL)
}

func downLoad(r Retriever) string{
	return r.Get("https://www.baidu.com/")
}



//接口的值类型

func main() {
	var r Retriever

	r = mock.Retriever{"this  is the values"}

	inspects(r)

	fmt.Println("\n要去执行del真实的方法啦~~~\n")
	r = &deal.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute,
	}
	inspects(r)

	realRetriever := r.(*deal.Retriever)

	fmt.Printf("\n\n细节：  %T \t %v\n\n",realRetriever,realRetriever)

	if mockRetrever ,ok := r.(mock.Retriever); ok{
		fmt.Printf("判断类型 ： contents : %s",mockRetrever.Contents)
	}else{
		fmt.Print("not is a right retorfit ")
	}


	dealRetriever := &deal.Retriever{}
	fmt.Printf("\n\n细节：  %T \t %v\n\n",dealRetriever,dealRetriever)

	fmt.Println("try  a session ~~~")

	var retrieverPoster RetrieverPoster
	retrieverPoster = mock.Retriever{
		Contents : URL,
	}
	fmt.Println("retrieverPoster~~~ ： ",session(retrieverPoster))

}
func inspects(retriever Retriever) {
	switch v := retriever.(type) {
	case mock.Retriever:
		result := retriever.Get("https://www.360.com")
		fmt.Println(result,downLoad(retriever))
		fmt.Printf("\n\n细节：  %T %v\n\n",retriever,retriever)
		fmt.Printf("类型 ： contents : %s",v.Contents)
	case *deal.Retriever:
		result := retriever.Get("https://www.baidu.com/")
		fmt.Println(result,downLoad(retriever))
		fmt.Printf("\n\n细节：  %T %v\n\n",retriever,retriever)
		fmt.Printf("类型 ：  : %d ， %s",v.TimeOut ,v.UserAgent)
	}
}





func main1() {
	//var r Retriever
	//r = mock.Retriever{
	//	"i am a retriever",
	//}
	//
 	//intss := 40
	//chars := "kaklt"
	//nodes :=  Node{
	//	Name :"张三",
	//	Age : 12,
	//	Weight :12.345,
	//}
	//ss = chars
	//ss = intss
	//ss = nodes
	//ss =  Node{
	//	Name :"张三",
	//	Age : 12,
	//	Weight :12.345,
	//}


	//r = deal.Retriever{"http://gc.ditu.aliyun.com/geocoding?a=%E8%8B%8F%E5%B7%9E%E5%B8%82",1000}
	//fmt.Println(downLoad(r)+"\n")

	// 定义a为空接口

	var i int = 5
	s := "Hello world"
	// a可以存储任意类型的数值
	a = i
	a = s
}
var a interface{

}
var ss interface {

}