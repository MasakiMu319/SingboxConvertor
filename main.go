package main

import (
	"SingboxConvertor/api"
	"SingboxConvertor/api/handlers"
	"SingboxConvertor/db"
	"SingboxConvertor/utils"
	"context"
	"github.com/gin-contrib/sessions"
	redisStore "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	authHandler *handlers.AuthHandler
	subHandler  *handlers.SubHandler
)

func init() {
	if err := db.InitMongoClient(); err != nil {
		utils.ConvertorLogPrintf(err, "Initialization of MongoDB client failed.")
		os.Exit(1)
	}
	collectionUsers := db.MongoClient.Database(db.DB).Collection(db.UserCollection)
	ctx := context.Background()

	authHandler = handlers.NewAuthHandler(ctx, collectionUsers)
	subHandler = &handlers.SubHandler{}
}

func main() {
	r := gin.Default()
	r.Static("/web", "./web")
	setupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_ = r.Run(":" + port)
}

func setupRoutes(router *gin.Engine) {
	store, _ := redisStore.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	store.Options(sessions.Options{
		MaxAge: 86400,
	})
	router.Use(sessions.Sessions("configurations_api", store))

	router.GET("/", api.GetFrontend)

	router.POST("/signin", authHandler.SignInHandler)
	router.POST("/signout", authHandler.SignOutHandler)
	router.POST("/signup", authHandler.SignUpHandler)
	router.POST("/refresh", authHandler.RefreshHandler)

	authorized := router.Group("/")
	authorized.Use(authHandler.AuthMiddleware())
	{
		authorized.GET("/sub", subHandler.GetSubHandler)
		authorized.POST("/generate", subHandler.GenSubHandler)
	}
}
