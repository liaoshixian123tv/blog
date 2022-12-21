package routers

import (
	v1 "blog/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
)

// BlogRouter 部落格路由
func BlogRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	var (
		tag    v1.Tag
		artcle v1.Artcle
	)

	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/artcles", artcle.Create)
		apiV1.DELETE("/artcles/:id", artcle.Delete)
		apiV1.PUT("/artcles/:id", artcle.Update)
		apiV1.PATCH("/artcles/:id/state", artcle.Update)
		apiV1.GET("/artcles/:id", artcle.Get)
		apiV1.GET("/artcles", artcle.List)
	}

	return router
}
