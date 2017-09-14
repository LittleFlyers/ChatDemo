package controllers

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
)

// Operations about Users
type ChartroomController struct {
	beego.Controller
}

/***新建聊天室****/
func (c *ChartroomController) Add() {
	c.SetSession("app", 1)
	db, err := sql.Open("mysql", "root:521wangjiaxuan.@/chart_room?charset=utf8")
	checkErr(err)
	id := c.GetString("id")
	name := c.GetString("name")
	introduction := c.GetString("introduction")
	//name := c.GetSession("app")
	stmt, err := db.Prepare(`INSERT chartroom (id,name,head,introduction) values (?,?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec(id, name, "null", introduction)
	checkErr(err)
	fmt.Println(res)
}
