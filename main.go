package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"itish.github.io/controller"
	"itish.github.io/initializers"
	"itish.github.io/service"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.ConnectCDB()
	initializers.SyncDB()
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	f, err := os.Create("logFile.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router.Static("/css", "./css")
	router.LoadHTMLGlob("html/*")

	r := router.Group("/v1")
	{
		r.GET("/home", service.HomePage)

		r.GET("/signup", service.SignUpPage)
		r.GET("/signupdone", service.SignUpDone)
		r.POST("/signup", controller.SignUp)

		r.GET("/login", service.LoginPage)
		r.GET("/logindone", service.LoginDone)
		r.POST("/login", controller.Login)

		r.GET("/blogpost", service.BlogPost)
		r.GET("/blogpostdone", service.BlogPostDone)
		r.GET("/blogpostedit", service.BlogEditPage)
		r.POST("/blogpost", controller.BlogCreate)
		r.POST("/blogpostedit", controller.BlogEdit)
	}
	router.Run()
}
