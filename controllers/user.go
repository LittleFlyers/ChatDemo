package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rongcloud/server-sdk-go/RCServerSDK"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

type Token struct {
	code         int
	token        string
	user_id      int
	errorMessage string
}

/******用户注册****/
func (u *UserController) SignIn() {
	//var saveToken Token
	user_id := u.GetString("user_id")
	user_password := u.GetString("user_password")
	rongcloud := rcserversdk.CreateRongCloud("lmxuhwagl00bd", "02FAas1bck")
	result, tokenError := rongcloud.User.GetToken(user_id, "1", "1")
	if result != nil || tokenError != nil {
		db, err := sql.Open("mysql", "root:521wangjiaxuan.@/chart_room?charset=utf8")
		checkErr(err)
		stmt, err := db.Prepare(`INSERT user (user_id, user_password, token) values (?,?,?)`)
		checkErr(err)
		res, err := stmt.Exec(user_id, user_password, result.Token)
		checkErr(err)
		fmt.Println(res)
	}
}

/*******用户登录*****/
func (u *UserController) LogIn() {
	user_id := u.GetString("user_id")
	user_password := u.GetString("user_password")
	db, err := sql.Open("mysql", "root:521wangjiaxuan.@/chart_room?charset=utf8")
	checkErr(err)
	rows, err := db.Query(`SELECT user_id,user_password,token FROM user WHERE user_id =` + user_id)
	tokenData := make(map[string]string)
	for rows.Next() {
		var user_id_check int
		var user_password_check string
		var token string
		rows.Scan(&user_id_check, &user_password_check, &token)
		if user_password == user_password_check {
			tokenData["token"] = token
			tokenData["msg"] = "success"
		} else {
			tokenData["msg"] = "password not corrcet"
		}
	}
	data, _ := json.Marshal(tokenData)
	u.Data["json"] = string(data)
	u.ServeJSON()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}
