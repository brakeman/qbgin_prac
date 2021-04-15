package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已经保存")
}
