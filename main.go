package main

import (
	"github.com/gin-gonic/gin"
	routes "github.com/weilinux/golang-jwt-project/routes"
	"os"
)

//对单机架构而言，这个项目的架构部署还是非常好的，很方便于扩展
//对单机架构而言，这个项目的架构部署还是非常好的，很方便于扩展


//side project for golang job
//side project for golang job
func main() {

	//dirty code dirty code dirty code
	//dirty code dirty code dirty code
	port := os.Getenv("PORT")
	if port==""{
		port="8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRouter(router)
	routes.UserRouter(router)
	//dirty code dirty code dirty code
	//dirty code dirty code dirty code


	router.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"Access granted for api-2"})
	})

	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}

}
