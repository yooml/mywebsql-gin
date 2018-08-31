package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"mywebsql-gin"
	//"log"
	//"html/template"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"time"
	"fmt"
)

//var engine *xorm.Engine
//var err error
//engine,err = xorm.NewEngine("mysql", "root:passwd@/test")

type User struct {
	Id int `json:"user_id"`
	User_code string `json:"user_code"`
	User_name string `json:"user_name"`
	User_pwd string `json:"user_pwd"`
	//Db_id int
	//User_id int
}

type Userdb struct {
	User_id int `json:"user_id"`
	Db_id int `json:"db_id"`
	End_datetime time.Time `json:"end_datetime"`
	Credite_datetime time.Time
	//Db_id int
	//User_id int
}


type Userdball struct {
	User_id int `json:"user_id"`
	Db_id int `json:"db_id"`
	End_datetime time.Time `json:"end_datetime"`
	Credite_datetime time.Time
	Db_name string `json:"db_name"`
	Db_host string `json:"db_host"`
	User_have bool `json:"user_have"`
	//User_id int
}

type Db_config struct {
	Id int
	Dbname string `form:"dbname" json:"dbname" binding:"required"`
	Dbhost string `form:"dbhost" json:"dbhost" binding:"required"`
	Dbuser string
	Dbpwd string
	Dbtype string
}

type Change_user struct {
	change_user_id string
}

type Userdbsql struct {
	User_id int `json:"user_id"`
	Db_id int `json:"db_id"`
	User_have bool `json:"user_have"`
	//Db_id int
	//User_id int
}


var engine *xorm.Engine

func InitDB() (*xorm.Engine,error){

	var err error
	db,err := xorm.NewEngine("mysql", "root:passwd@/test")
	engine=db
	return engine,err
}

func Sysuser(c *gin.Context) {
	c.HTML(http.StatusOK, "views/sysuser.html",nil)
	return

}

func AddSysuser(c *gin.Context)  {
	//var err error
	var user []User
	//var user User


	/*if err != nil {
		println("a",err.Error())
		return
	}*/
	user_code := c.PostForm("user_code")
	//user_pwd := c.PostForm("user_pwd")
	//user_name := c.PostForm("user_name")
	//s:="shen"
	sql_sel_user:="select * from sys_user where user_code like '%"+user_code+"%'"
	engine.Sql(sql_sel_user).Find(&user)
	for k,v :=range user{
		println(v.Id,v.User_code,v.User_name,k)
	}
	//println(user_code,user_name,user_pwd)
	c.HTML(http.StatusOK, "views/sysuser.html",nil)

}

func SelectSysuser(c *gin.Context)  {
	//var err error
	var user []User
	println("test")

	/*if err != nil {
		println("a",err.Error())
		return
	}*/
	select_user_code := c.PostForm("select")
	//user_pwd := c.PostForm("user_pwd")
	//user_name := c.PostForm("user_name")
	//s:="shen"
	sql_sel_user:="select * from sys_user where user_code like '%"+select_user_code+"%'"
	engine.Sql(sql_sel_user).Find(&user)
	users:=make([]User,0)
	for _,v :=range user{
		//var u User
		//println(v.Id,v.User_code,v.User_name,k)
		users=append(users,v)
	}
	//println(users[0].User_name)

	//println(user_code,user_name,user_pwd)
	/*c.JSON(http.StatusOK,gin.H{
		"data":users,
	})*/
	c.HTML(http.StatusOK, "views/sysuser.html",gin.H{
		"data":users,
	})

}


func Sysuserdb(c *gin.Context) {
	var userdb []Userdb
	var user_dbconfig []Db_config
	change_user_id := c.PostForm("cuserdb")
	println(change_user_id)
	sql_sel_user:="select * from sys_userdb where user_id ="+change_user_id
	engine.Sql(sql_sel_user).Find(&userdb)
	userdbs:=make([]Userdb,0)
	user_dbconfigs:=make([]Db_config,0)
	for _,v :=range userdb{
		userdbs=append(userdbs,v)
		sql_sel_db:="select * from db_config where id ="+fmt.Sprint(v.Db_id)
		engine.Sql(sql_sel_db).Find(&user_dbconfig)

		for _,v :=range user_dbconfig{
			user_dbconfigs=append(user_dbconfigs,v)
		}
	}


	/*var dbconfig []Db_config
	sel_dbname := c.PostForm("seldbname")
	sql_sel_db:="select * from db_config where dbname like '%"+sel_dbname+"%'"
	engine.Sql(sql_sel_db).Find(&dbconfig)
	sel_dbconfigs:=make([]Db_config,0)
	for _,v :=range dbconfig{
		sel_dbconfigs=append(sel_dbconfigs,v)
	}*/

	c.HTML(http.StatusOK, "views/sysuserdb.html",gin.H{
		"user_has_db":userdbs,
		"user_has_db_config":user_dbconfig,
	})
	//return

}



func SelectDb(c *gin.Context)  {
	var select_dbconfig []Db_config
	//c.ShouldBindJSON(&select_dbconfig)
	//sel_dbname:=select_dbconfig.Dbname
	//sel_dbhost:=select_dbconfig.Dbhost
	sel_dbname := c.PostForm("dbname")
	sel_dbhost := c.PostForm("dbhost")
	println(sel_dbname)
	if sel_dbhost !="" && sel_dbname =="" {
		sql_sel_dbhost:="select * from db_config where dbhost like '%"+sel_dbhost+"%'"
		engine.Sql(sql_sel_dbhost).Find(&select_dbconfig)
	}
	if sel_dbname !="" && sel_dbhost =="" {
		sql_sel_dbname:="select * from db_config where dbname like '%"+sel_dbname+"%'"
		engine.Sql(sql_sel_dbname).Find(&select_dbconfig)
	}
	if sel_dbhost !="" && sel_dbname !="" {
		sql_sel_dbname:="select * from db_config where dbname like '%"+sel_dbname+"%' and dbhost like '%"+sel_dbhost+"%'"
		engine.Sql(sql_sel_dbname).Find(&select_dbconfig)
	}

	//sel_dbs:=make([]Db_config,0)
	//for _,v :=range select_dbconfig{
	//	sel_dbs=append(sel_dbs,v)
	//}
	println(select_dbconfig[0].Dbname)
	//c.HTML(http.StatusOK, "views/sysuserdb.html",gin.H{
		//"user_has_db":sel_dbs,
		//"user_has_db_config":select_dbconfig,
	//})
	c.JSON(http.StatusOK, select_dbconfig)
}

func Testsql(c *gin.Context)  {
	var json Userdbsql
	if err := c.ShouldBindJSON(&json); err == nil {
		if json.User_have == true{
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in","user_id":json.User_id,"db_id":json.Db_id})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}


}

func Vuesysuser(c *gin.Context)  {
	var sysuser []User
	sql_sel_all_sysuser:="select * from sys_user"
	engine.Sql(sql_sel_all_sysuser).Find(&sysuser)
	c.JSON(http.StatusOK, sysuser)
}

func Vuesysuserdb(c *gin.Context)  {
	user_id := c.Query("user_id")
	//println(user_id)
	c.HTML(http.StatusOK, "views/vuesysuserdb.html",gin.H{
		"user_id":user_id,
	})
}

func Vuesysuserdbjson(c *gin.Context)  {
	var json Userdbsql
	c.ShouldBindJSON(&json)
	println(json.User_id)

	var userdball []Userdball
	var alluser_dbconfig []Db_config

	sql_sel_all_user:="select * from sys_userdb"
	engine.Sql(sql_sel_all_user).Find(&userdball)
	userdballs:=make([]Userdball,0)
	for _,vuserdball :=range userdball{
		vuserdball.User_have=false
		sql_sel_all_db:="select * from db_config where id ="+fmt.Sprint(vuserdball.Db_id)
		engine.Sql(sql_sel_all_db).Find(&alluser_dbconfig)

		for _,vdb :=range alluser_dbconfig{
			vuserdball.Db_name=vdb.Dbname
			vuserdball.Db_host=vdb.Dbhost
		}
		//vuserdba, _ := json.Marshal(vuserdb)
		userdballs=append(userdballs,vuserdball)
	}
	println(userdballs[0].Db_id,userdballs[0].Db_name)
	var userdb []Userdb
	//var user_dbconfig []Db_config

	sql_sel_user:="select * from sys_userdb where user_id ="+fmt.Sprint(json.User_id)
	engine.Sql(sql_sel_user).Find(&userdb)
	for _,vuserdb :=range userdb{
	for _,alluserdb:=range userdballs{
		if alluserdb.Db_id==vuserdb.Db_id{
			println(vuserdb.Db_id)
			alluserdb.User_have=true
			}else {
		continue
		}
	}
	}


	//userdb :=[{ user_id:467, db_id: 1,db_name:"finance_car_lease_v3",db_host:"bp184d696xe285rmlrw.mysql.rds.aliyuncs.com", user_have: true },]
	c.JSON(http.StatusOK, userdballs)

}
