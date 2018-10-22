package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const filePath2 = "E:\\englis_install\\go_pro\\src\\com.yuer.gio\\lgotest\\simple\\tempfile\\"

func main() {
	fmt.Println(filePath2)

	writeTiFile(filePath2 + "fmt2wl.yuer")
}
func writeTiFile(filePath string) {
	file, erro := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0660)

	//第一种方式
	if erro != nil {
		panic(erro)
	}
	defer file.Close()
	file.WriteString("12341234423")

	//第二种方式
	d1 := []byte(" 来 一 起 啪啪啪 Go")
	err := ioutil.WriteFile(filePath, d1, 0644)

	err = ioutil.WriteFile(filePath, []byte(" laaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"), 0644)
	checkError(err)

	//第三种方式
	w := bufio.NewWriter(file)
	w.Write(d1)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()

}
func checkError(erro error) {
	if erro != nil {
		panic(erro)
	}
}
