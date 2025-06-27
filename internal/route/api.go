package route

import (
	"go-web-template/internal/model/response"

	"github.com/gin-gonic/gin"
)

func setUpApi(r *gin.RouterGroup) {

	demo := r.Group("/demo")
	{
		demo.POST("/test", func(ctx *gin.Context) {
			response.Ok(ctx)
		})
	}
}
