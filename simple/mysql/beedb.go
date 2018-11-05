package main

import (
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
	"time"
)
func main() {
	db, err := sql.Open("mymysql", "test/root/yyh123")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)

	//user := tb_user{
	//	id:101,
	//	name:"贝吉塔",
	//	addr:"北京市朝阳区西湖灵隐寺",
	//	age:18,
	//	sex:"男",
	//	tel:"13011000000",
	//}

	//var  tbUser tbuser
	//tbUser.name = "贝吉塔"
	//tbUser.addr = "北京市朝阳区西湖灵隐寺"
	//tbUser.age = 18
	//tbUser.sex = "男"
	//tbUser.tel = "13011000000"
	//erro := orm.Save(&tbUser)

	var saveone Userinfo
	saveone.Username = "Test Add User"
	saveone.Departname = "Test Add Departname"
	saveone.Created = time.Now()
	erro := orm.Save(&saveone)


	if erro != nil{
		panic(erro)
	}
}
type tbuser struct {
	id int `PK` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	name string
	age int
	sex string
	addr string
	tel string
}

type Userinfo struct {
	Uid        int `PK` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username   string
	Departname string
	Created    time.Time
}