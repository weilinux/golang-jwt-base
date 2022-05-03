package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/weilinux/golang-jwt-project/controllers"
)

func UserRouter(incomingRoutes *gin.Engine) {
	//to ensure these routes are protected routes, 如何登录了系统，就会有token,所以需要下面的认证
	//he should not be able to access these users routes without user's token
	//如果整个系统变成很大的时候，也是可以通过这样限制的

	//incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())

}