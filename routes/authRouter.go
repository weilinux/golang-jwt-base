package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/weilinux/golang-jwt-project/controllers"
)


// 2.建立route mapping to function
// 3. dirty code dirty code dirty code
// 3. dirty code dirty code dirty code
func AuthRouter(incomingRoutes *gin.Engine)  {
	incomingRoutes.POST("users/signup", controller.SignUp())
	incomingRoutes.POST("users/login", controller.Login())
}