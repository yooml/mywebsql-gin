package main
import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"mywebsql-gin/controllers"
	"github.com/gin-gonic/gin/render"
)

type StructA struct {
	FieldA string `form:"field_a"`
}


type MyHTMLRender struct {
	templates map[string]*template.Template
}

func main() {
	router := gin.Default()
	//router.LoadHTMLGlob("views/*")
//router.LoadHTMLGlob("views/*.html")
//html := template.Must(template.ParseFiles("views/mywebsql-base.html","views/sysuser.html"))
//html := template.Must(template.ParseGlob("views/*"))
//router.SetHTMLTemplate(html)
	router.HTMLRender = createMyRender()
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

	router.GET("/test",func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/testvue.html",nil)
	})
	router.GET("/test2",func(c *gin.Context) {
		c.HTML(http.StatusOK, "test3",nil)
	})

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
