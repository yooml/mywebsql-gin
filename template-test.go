package main
import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"mywebsql-gin/controllers"
	"github.com/gin-gonic/gin/render"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

type StructA struct {
	FieldA string `form:"field_a"`
}


type MyHTMLRender struct {
	templates map[string]*template.Template
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
//router.LoadHTMLGlob("views/*.html")
//html := template.Must(template.ParseFiles("views/mywebsql-base.html","views/sysuser.html"))
//html := template.Must(template.ParseGlob("views/*"))
//router.SetHTMLTemplate(html)
	//router.HTMLRender = createMyRender()
	db,_ := controllers.InitDB()
	defer db.Close()
	router.Static("/static", "./static")
//router.LoadHTMLGlob("views/*.html")
	router.GET("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/login", controllers.Login)
	authorized := router.Group("/user", MyMiddelware())
	authorized.POST("/info", func(c *gin.Context) {
		c.String(http.StatusOK, "info")
	})
	authorized.GET("/info", func(c *gin.Context) {
		c.String(http.StatusOK, "info")
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/sysuser.html",nil)
	})
	router.POST("/addsysuser", controllers.AddSysuser)
	router.POST("/selectsysuser", controllers.SelectSysuser)
	router.POST("/selectdb", controllers.SelectDb)
	router.GET("/sysuser",controllers.Sysuser)
	router.POST("/sysuserdb",controllers.Sysuserdb)
	router.GET("/sysuserdb",controllers.Sysuserdb)
	router.POST("/vuesysuser",controllers.Vuesysuser)
	router.GET("/vuesysuserdb",controllers.Vuesysuserdb)
	router.POST("/vuesysuserdb",controllers.Vuesysuserdbjson)

	router.GET("/test",func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/testvuerouter.html",nil)
	})
	router.GET("/test2",func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/testvue2.html",gin.H{
			"a":"ll",
		})
	})

	authorized.GET("/vuesysuser",func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/vuesysuser.html",nil)
		})
	router.GET("/test4",func(c *gin.Context) {
		/*ll:=`[
		{ user_id:467, db_id: 1,db_name:"_v3",db_host:"bp184d696xe285rmlrw.mysql.rds.aliyuncs.com", user_have: true },
		{ user_id:467, db_id: 2,db_name:"lease",db_host:"bp184d696xe285rmlrw.mysql.rds.aliyuncs.com", user_have: true },
		{ user_id:467, db_id: 3,db_name:"finance_car",db_host:"bp184d696xe285rmlrw.mysql.rds.aliyuncs.com", user_have: true },
		{ user_id:467, db_id: 4,db_name:"finance",db_host:"bp184d696xe285rmlrw.mysql.rds.aliyuncs.com", user_have: true },
	]`*/
		c.HTML(http.StatusOK, "views/testvue3.html",gin.H{
			"a":"ll",
		})
	})



	router.POST("/testsql",controllers.Testsql)
	router.Run(":8080")
}


func (r *MyHTMLRender) Add(name string, html *template.Template) {
	if r.templates == nil {
		r.templates = make(map[string]*template.Template)
	}
	r.templates[name] = html
}

func (r *MyHTMLRender) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: r.templates[name],
		//Data:     data,
	}
}

func createMyRender() render.HTMLRender {
	r := &MyHTMLRender{}
	r.Add("test3", template.Must(template.ParseFiles("views/testvue2.html")))


	return r
}

func MyMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return []byte("mobile"), nil
			})
		if err == nil {
			if token.Valid {
				c.Next()
			} else {
				c.String(http.StatusUnauthorized, "Token is not valid")
			}
		} else {
			c.String(http.StatusUnauthorized, "Unauthorized access to this resource")
		}
	}
}

