package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
)

//https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/gin%E6%BA%90%E7%A0%81%E8%A7%A3%E8%AF%BB/gin%E7%89%9B%E9%80%BC%E7%9A%84context.html


//Param (我自己的叫法: 路由变量)
//gin的标准叫法是Parameters in path. restful风格api如/user/john, 这个路由在gin里面是/user/:name, 要获取john就需要使用Param函数
//name := c.Param("name")

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		//goland:noinspection ALL
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" && uid != userId {
		//goland:noinspection ALL
		err = errors.New("Unauthorized to access this resource !")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}



