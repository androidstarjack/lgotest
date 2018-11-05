package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)


func errorPanic(writer http.ResponseWriter, request *http.Request) error {
	panic("paint ")
}
type testtingUserError  string

func (e testtingUserError) Error() string{
	return e.Message()
}

func (e testtingUserError) Message() string{
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testtingUserError("userError")
}

func errNotfoundError(writer http.ResponseWriter, request *http.Request) error {
	return testtingUserError("errNotfoundError")
}

func errNoPerssionError(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}
func errUnKonwError(writer http.ResponseWriter, request *http.Request) error {
	return testtingUserError("errUnKonwError")
}

func errNoError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintf(writer,"请求正确，code : 200 ,no error")
	return nil
}


var test = []struct{
	h appHandler
	code int
	message string
}{
	{errorPanic,500, "Internal Server Error"},
	{errUserError,400, " User Server Error"},
	{errNotfoundError,404, "errNotfoundError"},
	{errNoPerssionError,405, " permission denied"},
	{errUnKonwError,500, " Internal Server Error"},
	{errNoError,200, " Internal Server Error"},
}

func TestErrorWrapper(t *testing.T) {
	//httptest.NewRequest()


	for _ ,tt := range test{
		f := errorWrap(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://localhost:8181/list/abcTest.txt",
			nil)

		f(response,request)
		content ,_ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(content),"\n")
		if content != nil{
			if response.Code != tt.code || body != tt.message{
				t.Errorf("except （%d ,%s） but acturely got （%d ,%s）",tt.code,tt.message, response.Code,body)
			}
		}
	}
}

func TestErrorWrapperInServer(t *testing.T){
	//for _ ,tt :=range test{
	//	f := errorWrap(tt.h)
	//	httptest.NewServer(nil)
	//}
}
