package initRouter

import (
	"ginqb/handler"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	index := router.Group("/")
	{
		//// 添加 Get 请求路由
		index.GET("", retHelloGinAndMethod)
	}

	// 添加 user
	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)
	}

	return router
}

func retHelloGinAndMethod(context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}
