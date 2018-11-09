package parse

import (
	"com.yuer.gio/lgotest/simple/crawler/concurrent/engine"
	"fmt"
	"regexp"
)
const cityListRegexp = `<a href="(http://www.zhenai.com/zhenghun/[a-z][A-z]+)"[^>]*>([^<]+)</a>`

func ParseCityListdata (data [] byte) engine.ParseResult{
	regexpp := regexp.MustCompile(cityListRegexp)
	result := regexpp.FindAllSubmatch(data,-1)
	parseRequest := engine.ParseResult{}
	limit := 10
	for _ ,val := range result{
		//	fmt.Printf("解析出第 %d 条数据 : %s\n",index,val )
		//	for _, match := range val{
		//fmt.Printf("城市：%s URL: %s \n",val[2],val[1] )
		//}
		parseRequest.Items  = append(parseRequest.Items,"city "+ string(val[2]))
		parseRequest.Requests = append(parseRequest.Requests,engine.Request{
			Url:string(val[1]),
			ParseFunc:ParseCity,
		})
		limit --
		//if limit == 0{
		//	break
		//}

	}
	fmt.Printf("长度：%d",len(result))
	return  parseRequest
}