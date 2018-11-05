package deal

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut  time.Duration
}

//返回一个结果
func (r *Retriever) Get(url string) string {
	resp , erro :=http.Get(url)
	if erro != nil{
		panic(erro)
	}
	request, resEorro  := httputil.DumpResponse(resp,true)
	if resEorro != nil{
		panic(request)
	}
	resp.Body.Close()
	return  string(request)
}