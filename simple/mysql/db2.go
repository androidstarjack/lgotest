package main

import (
	"com.yuer.gio/lgotest/simple/mysql/utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var userDetail = "CREATE TABLE IF NOT EXISTS userdetail    ( uid INT(10) NOT NULL, intro TEXT NULL, profile TEXT NULL, PRIMARY KEY (uid) ); "
var userInfo = "CREATE TABLE IF NOT EXISTS userinfo ( uid INTEGER PRIMARY KEY AUTO_INCREMENT,username VARCHAR(64) NULL,departname VARCHAR(64) NULL,created DATE NULL ); "

func main() {
	//创建数据库
	db, error := sql.Open("mysql", "root:yyh123@tcp(localhost:3306)/test?charset=utf8")
	if error != nil {
		panic(error)
	}

	_, err := db.Exec("use mysql")
	if err != nil {
		panic(err)
	}

	stmt, erro := db.Prepare(userDetail)
	if erro != nil {
		panic(erro)
	}

	result, err := stmt.Exec()
	if err != nil {
		panic(err)
	}

	affectId, err := result.RowsAffected()

	if err != nil {
		fmt.Printf("出错啦~  affectId :%d, %s", affectId, err)
	}

	fmt.Println("数据表：userDetail  成功！")

	stmt, err1 := db.Prepare(userInfo)
	if err1 != nil {
		panic(err1)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
	fmt.Println("数据表：userInfo 成功！")
	////表创建成功！！！！
	insertUserInfo(db)
	deleteTableItem(db)
}

func insertUserInfo(db *sql.DB) {
	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)
	res, err := stmt.Exec("astaxie", utils.GetNowtimeMD5(), "2012-12-09")
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Printf("插入的结果状态： %d", id)
}
func checkErr(e error) {
	if e != nil {
		fmt.Println("执行语句出错啦~~")
		panic(e)
	}
}

//删除数据’
func deleteTableItem(db *sql.DB) {
	stmt, err := db.Prepare("delete from  userinfo where uid = ?")
	checkErr(err)
	res, err := stmt.Exec(2)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Printf("删除的结果状态： %d", id)
}
