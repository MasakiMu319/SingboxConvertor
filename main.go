package main

import (
	"SingboxConvertor/internal/api"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./web")
	r.GET("/sub", api.GetSubscription)
	r.GET("/", api.GetFrontend)

	r.NoRoute(func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/")
	})

	// export PORT=8080
	if p := os.Getenv("PORT"); p != "" {
		_ = r.Run(":" + p)
	} else {
		_ = r.Run()
	}
}
