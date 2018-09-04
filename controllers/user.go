package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"mywebsql-gin"
	//"log"
	//"html/template"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/xormplus/xorm"
	"time"
	//"fmt"
	"github.com/dgrijalva/jwt-go"
	//"github.com/dgrijalva/jwt-go/request"
	"encoding/json"
)

type Myweb_user struct {
	//Id        int       `xorm:"pk autoincr"`
	Username  string    `xorm:"index unique notnull"`
	Password  string    `xorm:"password"`
	//CreatedAt time.Time `xorm:"created"`
}


func (c *Myweb_user) Register() string {
	engine.Insert(c)
	has,_ := engine.Get(c)
	if has {
		return JsonResponse(0, "注册成功")
	} else {
		return JsonResponse(1, "注册失败")
	}
}
func (c *Myweb_user) Login() string {
	has, _ := engine.Get(c)
	a:="q"
	if has {
		//println(&v)
		//return JsonResponse(0, setToken())
		return setToken()
	} else {
		return a
	}
}
func setToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte("mobile"))
	if err != nil {
		return ""
	}
	return tokenString
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := Myweb_user{
		Username: username,
		Password: password,
	}
	c.String(http.StatusOK, user.Register())
}
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	println(password)
	user := Myweb_user{
		Username: username,
		Password: password,
	}


	c.String(http.StatusOK, user.Login())
}

func JsonResponse(code int, data interface{}) string {
	response := make(map[string]interface{})
	response["code"] = code
	response["data"] = data
	js, err := json.Marshal(response)
	if err != nil {
		return err.Error()
	}
	return string(js)
}