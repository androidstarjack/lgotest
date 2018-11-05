package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r Retriever) String() string {
	 return fmt.Sprintln("为我们一起papapapapapapapapap","啪啪啪完之后，在睡觉")
}

func (r Retriever) post(url string,
	form map[string]string) string {
	panic("implement me")
}

func (r Retriever) papapa() {
	panic("implement me")
}

//func (r Retriever) papapa() {
//	panic("implement me")
//}

//func (r Retriever) post(url string,
//	form map[string]string) string {
//	panic("implement me")
//}

//func (Retriever) Get(url string) string {
//	panic("implement me")
//}


func (r Retriever) Get(url string) string{
	rresult := "从接口中实现的方法： " + url
	return rresult
}