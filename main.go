package main

import (
	"github.com/gin-gonic/gin"
	//"html/template"
	"net/http"
	"mywebsql-gin/controllers"
	//"path/filepath"
	//"log"
	"os"
	"io"
)

func main() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)
	router := gin.Default()
	//router.Static("/static", filepath.Join(getCurrentDirectory(), "./static"))
	router.Static("/static", "./static")
	router.LoadHTMLGlob("views/*")
	/*router.GET("/",  func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// use pusher.Push() to do server push
			if err := pusher.Push("/static/js/jquery.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
		//"title": "Main website",
		})
	})*/
	router.GET("/test",controllers.IndexGet)
	router.GET("/test2",controllers.Test2)
	//router.GET("/test",  func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tmpl",gin.H{
	//		"title": "Main website",
	//	})
	//})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/sysuser",controllers.Sysuser)
	router.Run()
}


/*func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		seelog.Critical(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}*/