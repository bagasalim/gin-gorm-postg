package main

import (
	"gin-gorm-postg/controllers"
	"gin-gorm-postg/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}
func main() {
	r := gin.Default()

	r.GET("/posts",controllers.Index) 
	r.GET("/posts/:id",controllers.Show) 
	r.POST("/posts",controllers.Create) 
	r.PUT("/posts/:id",controllers.Update) 
	r.DELETE("/posts/:id",controllers.Delete) 

	r.Run()
}