package router_v1

import (
	"web-api/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProcurement_SearchRouter(router *gin.RouterGroup) {
	router.GET("/procurement", controllers.Proc.GetPurListData)
	router.GET("/average", controllers.Proc.GetAverageDiffDay)
}
