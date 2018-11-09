package parse

import (
	"com.yuer.gio/lgotest/simple/crawler/simple/engine"
	"com.yuer.gio/lgotest/simple/crawler/simple/model"
	"regexp"
	"strconv"
)

const ProFileRuler  = `<td><span class="label">性别：</span><span field="">女</span></td>`
const ageProFileRuler  = `<td><span class="label">年龄：</span><span field="">35 - 48岁</span></td>`
const merryProFileRuler  = `<td><span class="label">婚况：</span>[^<]</td>`
const bodyHeightProFileRuler  = `<td><span class="label">身高：</span>^\\d+\.[0-9]+</td>`
const salaryProFileRuler  = `<td><span class="label">月收入：</span>[^<][^>]+</td>`
const educationProFileRuler  = `<td><span class="label">学历：</span>[^<][^>]+</td>`
const workAddressProFileRuler  = `<td><span class="label">工作地：</span>[[^<][^>]+</td>`
const prefressProFileRuler  = `<td><span class="label">职业： </span>[^<][^>]+</td>`
const isChildProFileRuler  = `<td><span class="label">有无孩子：</span>[^<][^>]+</td>`
const addreddProFileRuler  = `<td><span class="label">籍贯：</span>[^<][^>]+</td>`

var regexppAge = regexp.MustCompile(ageProFileRuler)
var merryProFileRulerData = regexp.MustCompile(merryProFileRuler)
var bodyHeightProFileRulerData = regexp.MustCompile(bodyHeightProFileRuler)
var salaryProFileRulerData = regexp.MustCompile(salaryProFileRuler)
var educationHeightProFileRulerData = regexp.MustCompile(educationProFileRuler)
var workAddressProFileRulerData = regexp.MustCompile(workAddressProFileRuler)
var prefressProFileRulerData = regexp.MustCompile(prefressProFileRuler)
var isChildProFileRulerData = regexp.MustCompile(isChildProFileRuler)
var addreddProFileRulerData = regexp.MustCompile(addreddProFileRuler)

func parseProFile(contents [] byte) engine.ParseResult{

	//macthData := merryProFileRulerData.FindSubmatch(contents)
	profile := model.Profile{

	}

		age ,erro := strconv.Atoi(extraRegexpString(contents,regexppAge))
		if erro != nil{
			panic(erro)
		}
		profile.Age = age


		merry:=extraRegexpString(contents,merryProFileRulerData)
		profile.Marriage = string(merry[1])

		bodyHeight :=extraRegexpInt(contents,bodyHeightProFileRulerData)
		profile.Height =bodyHeight

		address :=extraRegexpString(contents,workAddressProFileRulerData)
	 	profile.House = address

		return engine.ParseResult{
			Items:[]interface{}{profile},
		}
}

func extraRegexpString(contents [] byte,regespRuler * regexp.Regexp)  string{
	match := regespRuler.FindSubmatch(contents)
	if match != nil && len(match) >=2{
		return  string(match[1])
	}
	return ""
}
func extraRegexpInt(contents [] byte,regespRuler * regexp.Regexp)  int{
	match := regespRuler.FindSubmatch(contents)
	if match != nil && len(match) >=2{
		result, erro :=  strconv.Atoi(string(match[1]))
		if erro != nil{
			return 0
		}
		return result
	}
	return 0
}