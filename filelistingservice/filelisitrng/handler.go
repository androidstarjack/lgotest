package filelistingservice

import (

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)
const preFix = "/list/"

type userError  string

func (e userError) Error() string{
	return e.Message()
}

func (e userError) Message() string{
	return string(e)
}

func HandelFileListring(writer http.ResponseWriter, request *http.Request) error{
	index := strings.Index(request.URL.Path,preFix)
	log.Printf("ssssssssssssssssss: %d",index)
	if index !=0 {
		return userError("path must start with" +
			preFix+" or occurred a error")
	}
	path := request.URL.Path[len(preFix):]
	fmt.Println(path)

	file , err := os.Open(path)
	if err != nil{
		//panic(err)
		//http.Error(writer,
		//	err.Error(),
		//	http.StatusInternalServerError)
		return err
	}
	defer  file.Close()

	//然后是从文件里面去读内容
	content,readErr := ioutil.ReadAll(file)
	if readErr != nil{
		//panic(readErr)
		return readErr
	}
	//writer.Write([]byte(`hello world`))
	writer.Write(content)
	return nil
}
