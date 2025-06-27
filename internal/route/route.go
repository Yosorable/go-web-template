package route

import (
	"fmt"
	"io/fs"
	"math/rand"
	"net/http"

	"go-web-template/assets"
	"go-web-template/internal/controller"
	"go-web-template/internal/global"
	"go-web-template/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CreateRoute() (*gin.Engine, error) {
	fSys, err := fs.Sub(assets.WebStaticFiles, "web")
	if err != nil {
		return nil, err
	}

	if global.CONFIG.Secret == "" {
		global.CONFIG.Secret = fmt.Sprintf("%v", rand.Float64())
	}

	r := gin.Default()

	r.GET("/*any", func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		ctx.FileFromFS(path, http.FS(fSys))
	})

	auth := r.Group("/auth")
	{
		ctr := controller.AuthController
		auth.POST("/login", ctr.Login)
		auth.POST("/user", ctr.User)
	}

	r.Use(middleware.JWTAuthMiddleware())

	setUpApi(r.Group("/api"))

	return r, nil
}
