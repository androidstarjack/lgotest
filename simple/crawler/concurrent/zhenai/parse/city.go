package parse

import (
	"com.yuer.gio/lgotest/simple/crawler/concurrent/engine"
	"regexp"
)
                  //<a href="http://album.zhenai.com/u/1572218980" target="_blank">安然</a>
const cityRegexp =`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
func ParseCity(cotnent [] byte) engine.ParseResult{
	regexpp := regexp.MustCompile(cityRegexp)
	result := regexpp.FindAllSubmatch(cotnent,-1)
	parseRequest := engine.ParseResult{}
	//fmt.Printf("ParseCity： %s",result)
	for _ ,val := range result{

		parseRequest.Items  = append(parseRequest.Items, "user: "+string(val[2])) //用户的名字
		parseRequest.Requests = append(parseRequest.Requests,engine.Request{
			Url:string(val[1]),
			ParseFunc:engine.NilParseReuqest,
		})

	}
	//fmt.Printf("长度：%d",len(result))
	return parseRequest
}