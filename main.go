package main

import (
	"SingboxConvertor/internal/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods:     []string{"GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Static("/web", "./web")

	r.GET("/sub", api.GetSubscription)
	r.POST("/generate", api.PostGenerate)
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
