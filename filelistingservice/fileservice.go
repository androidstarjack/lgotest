package main

import (
	_ "net/http/pprof"
	"com.yuer.gio/lgotest/filelistingservice/filelisitrng"
	"fmt"
	"log"
	"net/http"
	"os"
)


type appHandler func (writer http.ResponseWriter, request *http.Request) error

func errorWrap(handler appHandler) func(http.ResponseWriter,   *http.Request){
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r!= nil{
				log.Printf("panic :%v \n",r)
				fmt.Println("报错啦~报错啦报错啦报错啦报错啦报错啦")
				fmt.Printf("xxxx %v",r)
				http.Error(responseWriter,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			}
		}()
		err := handler(responseWriter,request)
		if err != nil{

			log.Printf("Error occurred  handler request %v",err)

			if userErr , ok := err.(userError); !ok{
				log.Println("我怕怕啪啪啪啪啪啪怕",userErr.Message())
				http.Error(responseWriter,userErr.Message(),http.StatusInternalServerError)
				return
			}
			//fmt.Println("报错啦",err.Error())
			if err != nil{
				fmt.Println("报错啦",err.Error())
				code := http.StatusOK
				switch   {
				case os.IsNotExist(err):
					code = http.StatusNotFound
				case os.IsPermission(err):
					code = http.StatusForbidden
				default:
					code = http.StatusInternalServerError
				}

				http.Error(responseWriter,http.StatusText(code),code)
			}else {
				fmt.Println("啦啦啦，我正常进行中.....")
			}
		}

	}
}

type userError interface {
	error
	Message() string
}
func main() {
	http.HandleFunc("/list/",errorWrap(filelistingservice.HandelFileListring))
	err := 	http.ListenAndServe(":8181",nil)
	if err != nil{
		panic(err)
	}
}
