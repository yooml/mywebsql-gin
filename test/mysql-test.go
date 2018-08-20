package main

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/xormplus/xorm"
    //"fmt"
    //"database/sql"
)

//var engine *xorm.Engine




func main() {
    var err error
    var user []User
    //var user User
    engine, err = xorm.NewEngine("mysql", "root:GTRawfQKiGmVIoFXSlcgdnlvZOXSO8@/test")
    //engine, err = xorm.NewEngine("mysql", "root:GTRawfQKiGmVIoFXSlcgdnlvZOXSO8@tcp(47.98.173.194:13306)/test")
    //fmt.Printf(err.Error())

    if err != nil {
        println("a",err.Error())
        return
    }
    println("sss")
    //sql_sel_userdb := "select * from sys_userdb where user_id=?"
    s:="shen"
    sql_sel_user:="select * from sys_user where user_code like '%"+s+"%'"
    engine.Sql(sql_sel_user).Find(&user)
    //println(err.Error())
    //record := make(xorm.Record)
    //has, err = session.SQL("select * from sys_userdb where id=?", 2).Get(&record)
    //id := record["id"].Int64()
    for k,v :=range user{
        println(v.Id,v.User_code,v.User_name,k)
    }
    //println(user[0].Db_id,user[0].Id)
}

