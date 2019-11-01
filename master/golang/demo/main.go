package main

import (
	"fmt"
	"github.com/gohouse/gorose/v2"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Uid int64 `gorose:"uid"`
	Name string `gorose:"name"`
	Age int64 `gorose:"age"`
	Xxx interface{} `gorose:"-"` // 这个字段在orm中会忽略
}

func (u *User) TableName() string {
	return "users"
}

var err error
var engin *gorose.Engin

func init() {
	// 全局初始化数据库,并复用
	// 这里的engin需要全局保存,可以用全局变量,也可以用单例
	// 配置&gorose.Config{}是单一数据库配置
	// 如果配置读写分离集群,则使用&gorose.ConfigCluster{}
	engin, err = gorose.Open(&gorose.Config{Driver: "sqlite3", Dsn: "./db.sqlite"})
}
func DB() gorose.IOrm {
	return engin.NewOrm()
}
func main() {
	// 这里定义一个变量db, 是为了复用db对象, 可以在最后使用 db.LastSql() 获取最后执行的sql
	// 如果不复用 db, 而是直接使用 DB(), 则会新建一个orm对象, 每一次都是全新的对象
	// 所以复用 db, 一定要在当前会话周期内
	db := DB()

	// 查询一条
	var u User
	// 查询数据并绑定到 user{} 上
	err = db.Table(&u).Fields("uid,name,age").Where("age",">",0).OrderBy("uid desc").Select()
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(u, u.Name)
	fmt.Println(db.LastSql())

	// 查询多条
	// 查询数据并绑定到 []Users 上, 这里复用了 db 及上下文条件参数
	// 如果不想复用,则可以使用DB()就会开启全新会话,或者使用db.Reset()
	// db.Reset()只会清除上下文参数干扰,不会更换链接,DB()则会更换链接
	var u2 []User
	err = DB().Table(&u2).Limit(10).Offset(1).Select()
	fmt.Println(u2)
	fmt.Println(db.LastSql())

	// 统计数据
	var count int64
	// 这里reset清除上边查询的参数干扰, 可以统计所有数据, 如果不清楚, 则条件为上边查询的条件
	// 同时, 可以新调用 DB(), 也不会产生干扰
	count,err = db.Reset().Count()
	// 或
	//count, err = DB().Table(&u).Count()
	fmt.Println(count, err)
	fmt.Println(db.LastSql())
}