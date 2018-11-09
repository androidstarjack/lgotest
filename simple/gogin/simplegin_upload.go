package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	log.Print("-----------go gin ----------------")
	requestMain()
}


type  Name  func(context *gin.Context)

func fiefie(c *gin.Context) {
	name := c.Param("name")
	passw := c.Param("password")

	if name == "yuer" && passw == "123456" {
		fmt.Printf("name: %s   password:  %s",name,passw)
	}else{
		log.Print("-----------密码和账号错误----------------")
	}
}

func requestMain(){
	router := gin.Default()

	router.GET("/list", func(c *gin.Context) {
		log.Print("-----------   end  ---------------")
		name := c.DefaultQuery("name","root")
		passw := c.DefaultQuery("password","0")

		if name == "yuer" && passw == "123456" {
			fmt.Printf("name: %s   password:  %s",name,passw)
			c.JSON(http.StatusOK,gin.H{
				"status_code": http.StatusOK,
				"message":"恭喜您，可以一块啪啪啪了",
				"token":"afds1345adsfhasdbv[zxcu adsvzxhcvzxcnv;aufvalsdvalsydvc",
			})
		}else{
			log.Print("-----------密码和账号错误----------------")
			c.JSON(http.StatusOK,gin.H{
				"status_code": http.StatusOK,
				"message":"用户名名或者密码错误~",
				"token":nil,
			})
		}
	})
	//manners.ListenAndServe(":8888", router)
	router.Run(":8888")
}