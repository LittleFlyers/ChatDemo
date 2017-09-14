package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
)

// Operations about Users
type ChartController struct {
	beego.Controller
}

/****聊天*****/
func (c *ChartController) Add() {
	c.SetSession("app", 1)
	db, err := sql.Open("mysql", "root:521wangjiaxuan.@/chart_room?charset=utf8")
	checkErr(err)
	from_user := c.GetString("from_user")
	to_user := c.GetString("to_user")
	message := c.GetString("message")
	send_time := time.Now().Format("2006-01-02 15:04:05")
	state := 1
	stmt, err := db.Prepare(`INSERT INTO charts (from_user,message,to_user,state,send_time) values (?,?,?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec(from_user, message, to_user, state, send_time)
	checkErr(err)
	fmt.Println(res)
	fmt.Println(send_time)
}

/*****获取聊天内容****/
func (c *ChartController) GetChat() {
	c.SetSession("app", 1)
	db, err := sql.Open("mysql", "root:521wangjiaxuan.@/chart_room?charset=utf8")
	checkErr(err)
	to_user := c.GetString("to_user")
	rows, err := db.Query(`SELECT from_user,message,send_time FROM chart_room.charts where to_user = ` + to_user + ` ORDER BY send_time`)
	chatData := make(map[int]map[string]string)
	i := 0
	for rows.Next() {

		chatTemp := make(map[string]string)
		var message string
		var send_time string
		var from_user string
		err = rows.Scan(&from_user, &message, &send_time)
		chatTemp["from_user"] = from_user
		chatTemp["message"] = message
		chatTemp["send_time"] = send_time
		chatData[i] = chatTemp
		i++
	}

	data, _ := json.Marshal(chatData)
	fmt.Println(string(data))
	c.Data["json"] = string(data)
	c.ServeJSON()
}
