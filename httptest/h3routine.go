package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
)
type MyMux struct {
}
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.URL.Path:  ",r.URL.Path)
	if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		fmt.Println("未匹配到~~~~~")
		w.Write([]byte("未匹配到~~~~~"))
		return
	}
	if strings.Contains(r.URL.Path,"name") {
	//if r.URL.Path == "/" {
		sayhelloName2(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func sayhelloName2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}
func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
	t := time.Date(2009, time.November, 10, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
}