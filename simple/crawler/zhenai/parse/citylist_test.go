package parse

import (
	"com.yuer.gio/lgotest/simple/crawler/fecther"
	"fmt"
	"testing"
)

func TestParseCityListdata(t *testing.T) {

	data ,erro := fecther.ParseUrlByUtf8Tansfer("http://www.zhenai.com/zhenghun")
	if erro != nil{
		panic(erro)
	}
	fmt.Println(string(data))
	result := ParseCityListdata(data)
	const resultSize = 408
	var exceptUrls = [] string {
		"http://www.zhenai.com/zhenghun/dalian",
		"http://www.zhenai.com/zhenghun/dandong",
		"http://www.zhenai.com/zhenghun/daqing",
	}
	var exceptCityNames = [] string{
		"city 大连",
		"city 丹东",
		"city 大庆",
	}
	for index , url := range  exceptUrls{
		//if exceptUrls[index] != url{
		//		t.Errorf("except Url :#%d: %s ;but was  %s",index,url,  exceptUrls[index])
		//}
		if result.Requests[index].Url != url{
			t.Errorf("except Url :#%d: %s ;but was  %s",index,url,  exceptUrls[index])
		}
	}
	for index , cityName := range  exceptCityNames{
		//if exceptCityNames[index] != cityName{
		//	t.Errorf("except city :#%d: %s ;but was  %s",index,cityName,  exceptCityNames[index])
		//}
		if result.Items[index].(string) != cityName{
			t.Errorf("except city :#%d: %s ;but was  %s",index,cityName,  exceptCityNames[index])
		}
	}
	//<a href="http://www.zhenai.com/zhenghun/dalian" data-v-0c63b635>大连</a>
	//	<a href="http://www.zhenai.com/zhenghun/dandong" data-v-0c63b635>丹东</a>
	//		<a href="http://www.zhenai.com/zhenghun/daqing" data-v-0c63b635>大庆</a>
	if len(result.Requests) != resultSize{
		t.Errorf("result should has %d Reuqests ,but had %d",resultSize,len(result.Requests))
	}

	if len(result.Items) != resultSize{
		t.Errorf("Items should has %d Reuqests ,but had %d",resultSize,len(result.Requests))
	}
}

