package main

import "fmt"

const myname = "<a href=\"http://www.zhenai.com/zhenghun/nujiang\" data-v-0c63b635=\"\">怒江</a>"
const myname2 = `<a href="http://www.zhenai.com/zhenghun/nujiang" data-v-0c63b635="">怒江</a> <a href="http://www.zhenai.com/zhenghun/zhoukou" data-v-0c63b635="">商水</a>`

func main() {
/*	regp := regexp.MustCompile( `<a href="http://www.zhenai.com/zhenghun/[a-z][A-z]*"\s*data-v-0c63b635=""`)
	regp  = regexp.MustCompile( `<a href="http://www.zhenai.com/zhenghun/[a-z][A-z]+"[^>]*>[^<]+</a>`)
	//regp  = regexp.MustCompile( `<a href="http://www.zhenai.com/zhenghun/[^>]*>`)
	//regp := regexp.MustCompile( `<a href="http://www.zhenai.com/zhenghun/[a-z][A-z]+"[^>]*[a-z][A-Z]+[^<]</a>`)
	data := regp.FindAllString(myname2,-1)

	//if erro != nil{
	//	panic(erro)
	//}
	fmt.Println(data)*/
	array :=[5]int{0,1,2,3,4}
	for len(array)>0 {
		fmt.Println(array)
	}
}
