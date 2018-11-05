package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GetTime() string {
	const shortForm = "2006-01-02 15:04:05"
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	fmt.Println(t)
	return str
}

func GetMD5Hash(text string) string {
	haser := md5.New()
	haser.Write([]byte(text))
	return hex.EncodeToString(haser.Sum(nil))
}

func GetNowtimeMD5() string {
	t := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	return GetMD5Hash(timestamp)
}


func Generate_Randnum() int {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(100)

	fmt.Printf("rand is %v\n", rnd)

	return rnd
}