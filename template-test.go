package main
import (
	"github.com/gin-gonic/gin"
	//"html/template"
	"net/http"
	"mywebsql-gin/controllers"
)

type StructA struct {
	FieldA string `form:"field_a"`
}


func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
//router.LoadHTMLGlob("views/*.html")
//html := template.Must(template.ParseFiles("views/mywebsql-base.html","views/sysuser.html"))
//html := template.Must(template.ParseGlob("views/*"))
//router.SetHTMLTemplate(html)
	db,_ := controllers.InitDB()
	defer db.Close()
	router.Static("/static", "./static")
//router.LoadHTMLGlob("views/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/sysuser.html",nil)
	})
	router.POST("/addsysuser", controllers.AddSysuser)
	router.POST("/selectsysuser", controllers.SelectSysuser)
	router.POST("/selectdb", controllers.SelectDb)
	router.GET("/sysuser",controllers.Sysuser)
	router.POST("/sysuserdb",controllers.Sysuserdb)
	router.GET("/sysuserdb",controllers.Sysuserdb)

//router.GET("/test",IndexApi)

	router.Run(":8080")
}

