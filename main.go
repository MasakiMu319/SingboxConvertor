package main

import (
	"SingboxConvertor/internal/api"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()

	// TODO: get external configuration should be an inner op.
	//r.GET("/config", api.GetConfig)
	r.GET("/sub", api.GetSubscription)
	r.GET("/", api.GetFrontend)

	// export PORT=8080
	if p := os.Getenv("PORT"); p != "" {
		_ = r.Run(":" + p)
	} else {
		_ = r.Run()
	}
}
