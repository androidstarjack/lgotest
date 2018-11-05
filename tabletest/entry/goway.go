package entry

import (
	"bytes"
	"strings"
)

func checkHasSum (a ,b string) string {
	var stringbulder  strings.Builder
	stringbulder.WriteString(a)
	stringbulder.WriteString(b)
	return  stringbulder.String()
}


func chechhasSum2 (a ,b string) string {
	var buffer bytes.Buffer
	buffer.WriteString(a)
	buffer.WriteString(b)
	return  buffer.String()
}


func chechhasSum3(a ,b string) func(a,b string) {

	return func(a, b string) {
		var buffer bytes.Buffer
		buffer.WriteString(a)
		buffer.WriteString(b)
	}
}

