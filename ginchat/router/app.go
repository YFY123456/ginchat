package router

import (
	"ginchat/docs"
	"ginchat/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/", service.GetIndex) // 添加这行来处理根URL
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", service.GetIndex)
	r.GET("/toRegister", service.ToRegister)
	r.GET("/toChat", service.ToChat)
	r.GET("/chat", service.Chat)
	r.POST("/searchFriends", service.SearchFriends)
	r.POST("/user/createUser", service.CreateUser)
	r.GET("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)
	r.GET("/user/sendMsg", service.SendMsg)
	r.GET("/user/sendUserMsg", service.SendUserMsg)
	r.POST("/attach/upload", service.Upload)
	//静态资源
	r.Static("/asset", "asset/")
	r.LoadHTMLGlob("views/**/*")
	r.POST("/contact/addfriend", service.AddFriend)
	//创建群
	r.POST("/contact/createCommunity", service.CreateCommunity)
	//群列表
	r.POST("/contact/loadcommunity", service.LoadCommunity)
	r.POST("/contact/joinGroup", service.JoinGroups)
	r.POST("/user/find", service.FindByID)
	r.POST("/user/redisMsg", service.RedisMsg)
	return r
}
