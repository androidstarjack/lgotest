package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestTb(t *testing.T) {
	test := []struct {
		a, b, result string
		isReal       bool
	}{
		{"kaka", "1", "kaka1", true},
		{"卡卡", "罗特", "卡卡罗特", false},
		{"于", "亚豪", "豪", true},
		{"百", "年好合", "百年好和", false},
	}

	for _, entry := range test {
		var stringbuilder strings.Builder
		stringbuilder.WriteString(entry.a)
		stringbuilder.WriteString(entry.b)

		if entry.result == stringbuilder.String() {
			entry.isReal = true
		} else {
			entry.isReal = false
		}
		if !entry.isReal {
			fmt.Printf("出错条目： 字段：%s  字段：%s 原来结果：%s 正确结果: %s\n", entry.a, entry.b, entry.result, stringbuilder.String())
		}
	}
}

func Benchmark(tb *testing.B) {
	test := struct {
		a, b, result string
		isReal       bool
	}{"百", "年好合", "百年好和", false}
	for i := 0; i < tb.N; i++ {
		var stringbuilder strings.Builder
		stringbuilder.WriteString(test.a)
		stringbuilder.WriteString(test.b)

		if test.result == stringbuilder.String() {
			test.isReal = true
		} else {
			test.isReal = false
		}
		if !test.isReal {
			tb.Logf("出错条目： 字段：%s  字段：%s 原来结果：%s 正确结果: %s\n", test.a, test.b, test.result, stringbuilder.String())
		}
	}
	//原来如此
}
