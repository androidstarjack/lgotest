
package main

import (
	"com.yuer.gio/lgotest/simple/mysql/utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func insert(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO user(username, password) VALUES(?, ?)")
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return
	}
	stmt.Exec("guotie", "guotie")
	stmt.Exec("testuser", "123123")

}
var CURRENT_AGE = 20
var  sex   = "男"

func checkError(str  string,err error) bool{
	if err !=  nil{
		fmt.Printf(str+" %s \b \n",err.Error())
		panic(err)
		return false
	}
	return true
}

func main() {
	db, err := sql.Open("mysql", "root:yyh123@tcp(localhost:3306)/test?charset=utf8")
	checkError("打开一个 数据库",err)

	//创建数据库
	//createDataBase(db)
	//userDb(db)

	createTable(db)
	insertTableContent(db)
	queryFromDb(db)
	updataFromDb(db)
	//deleteFromTabCase(db)
	//dropTab(db)
}
func dropTab(db *sql.DB) {
	res ,erro := db.Exec("drop table tb_user")
	if erro != nil{
		panic(erro)
	}
	 affect,erro := res.RowsAffected()
	 if erro != nil{
		checkError("删除表\n",erro)
	}
	fmt.Printf("\n删除表成功 ，结果影响的行数是:%d\n",affect)
}
func deleteFromTabCase(db *sql.DB) {
	 stmt ,err := db.Prepare("delete from tb_user where name = ?")
	 checkError("根据条件进行删除表",err)

	res ,err :=  stmt.Exec("卡卡罗特")
	if err!= nil{
		panic(err)
	}

	affect, err := res.RowsAffected()
	lastId, err := res.LastInsertId()
	fmt.Printf("affect : %d   lasetId: %d",affect,lastId)
}
func updataFromDb(db *sql.DB) {
	stmt ,err := db.Prepare("update tb_user set name = ? where name = ?")
	checkError("查询条件数据库",err)
	result ,erro := stmt.Exec("卡卡罗特","yuer")
	if erro != nil{
		checkError("查询条件数据库",erro)
	}
	affect ,err := result.RowsAffected()
	checkError("查询的结果",err)
	fmt.Printf("更新的数据：%d",affect)
}

func queryFromDb(db *sql.DB) {
	row, error := db.Query("select * from tb_user")

	 if checkError("查询数据库",error){
	 	defer  row.Close()
	 		for row.Next(){
	 			var id int
	 			var name string
	 			var age int
	 			var sex string
	 			var addr string
	 			var tel string
	 			row.Scan(&id,&name,&age,&sex,&addr,&tel)

	 			fmt.Printf("查询到了： id：  %d  %s %d  %s  %s  %s \n",id,name,age ,sex,addr,tel)
			}


	 }
	//row, error := db.Query("select * from tb_user")
}




func insertTableContent(db *sql.DB) {
	//var userId int = utils.GetNowtimeMD5()
	stmt ,err := db.Prepare("insert   tb_user set id = ?, name = ? ,age = ?, sex = ?,addr = ?,tel=?;")
	//stmt, err := db.Prepare("insert userinfo set username=?,departname=?,created=?,password=?,uid=?")
	checkError("准备阶段，回准备要执行的sql操作，然后返回准备完毕的执行状态。",err)

	if CURRENT_AGE % 2 == 0{
		sex = "男"
	}else{
		sex = "女 "
	}
	CURRENT_AGE = CURRENT_AGE+utils.Generate_Randnum()
	result, err :=stmt.Exec(CURRENT_AGE,"yuer",CURRENT_AGE,sex,"河南省商水县等城镇林村","13011007869")
	if err != nil{
		panic(err)
	}
	fmt.Println("插入数据成功",result)
}
func createTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS tb_user(id int(10) primary key,name varchar(20),age int(10),sex varchar(5),addr varchar(64),tel varchar(11));")
	if err != nil {
		fmt.Println("create table  failed:", err.Error())
		return
	}
	fmt.Println("创建表成功啦~~")
}

func createDataBase(db *sql.DB) bool{
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS my_db;")
	return checkError("试图创建一个数据库",err)
}


//func userDb(db *sql.DB){
//	_, err := db.Exec("USE my_db;")
//	if err != nil {
//		fmt.Println("select database failed")
//		return
//	}
//}







func main1() {
	db, err := sql.Open("mysql", "root:yyh123@tcp(10.2.0.215:3306)/test?charset=utf8")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	insert(db)

	rows, err := db.Query("select id, username from user where id = ?", 1)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()
	var id int
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}