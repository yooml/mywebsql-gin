package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"mywebsql-gin"
	"log"
	//"html/template"
)

//var templates = template.Must(template.ParseGlob("views/*"))



func IndexGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
	//user, _ := c.Get(CONTEXT_USER_KEY)
	//c.HTML(http.StatusOK, "index/index.html")
}

func Test2(c *gin.Context) {
	if pusher := c.Writer.Pusher(); pusher != nil {
		// use pusher.Push() to do server push
		if err := pusher.Push("/static/js/jquery.js", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}


	c.HTML(http.StatusOK, "index.html", gin.H{
	//"title": "Main website",
	})
}



