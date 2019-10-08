package route

import (
	"github.com/gin-gonic/gin"
	"mobileSign/api"
	"mobileSign/middleware"
)

func Load(router *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// 入口
	router.GET("/user/enter", api.Enter)

	router.Use(middleware.AuthMiddleware())
	router.Use(middleware.CorsMiddleware())

	QAGroup := router.Group("/QA")
	{
		IntergralGroup := QAGroup.Group("/Intergral")
		{
			IntergralGroup.POST("/checkIn", api.Sign)
		}

		UserGroup := QAGroup.Group("/User")
		{
			UserGroup.GET("/getScoreStatus", api.SignInfo)
		}

	}
	return router
}
