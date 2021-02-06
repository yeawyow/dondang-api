package api

import (
	"main/db"
	"main/interceptor"
	"main/model"

	"github.com/gin-gonic/gin"
)

func setupUserAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	{

		authenAPI.POST("/edituser", edituser)
		authenAPI.POST("/getuser", interceptor.JwtVerify, getuser)
	}
}

func getuser(c *gin.Context) {
	var user []model.Users
	db.GetDB().Find(&user)
	c.JSON(200, gin.H{"data": user})
}

func edituser(c *gin.Context) {
	c.JSON(200, gin.H{"stus": "ok"})
}
