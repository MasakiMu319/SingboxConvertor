package handlers

import (
	"SingboxConvertor/api/model"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type AuthHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAuthHandler(ctx context.Context, collection *mongo.Collection) *AuthHandler {
	return &AuthHandler{
		collection: collection,
		ctx:        ctx,
	}
}

func (handler *AuthHandler) SignInHandler(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h := sha256.New()
	cur := handler.collection.FindOne(handler.ctx, bson.M{
		"username": user.Username,
		"password": hex.EncodeToString(h.Sum([]byte(user.Password))),
	})
	if cur.Err() != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	sessionToken := xid.New().String()
	session := sessions.Default(ctx)
	session.Set("username", user.Username)
	session.Set("token", sessionToken)
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User signed in",
	})
}

func (handler *AuthHandler) SignOutHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Signed out...",
	})
}

func (handler *AuthHandler) SignUpHandler(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check for duplicate registrations.
	filter := bson.D{{"username", user.Username}}
	var result bson.M
	err := handler.collection.FindOne(handler.ctx, filter).Decode(result)
	if !errors.Is(err, mongo.ErrNoDocuments) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": "username already existed",
		})
		return
	}

	h := sha256.New()
	_, err = handler.collection.InsertOne(handler.ctx, bson.M{
		"username": user.Username,
		"password": hex.EncodeToString(h.Sum([]byte(user.Password))),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Signed up successfully",
	})
}

func (handler *AuthHandler) RefreshHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	sessionToken := session.Get("token")
	sessionUser := session.Get("username")
	if sessionToken == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid session",
		})
		return
	}

	sessionToken = xid.New().String()
	session.Set("username", sessionUser.(string))
	session.Set("token", sessionToken)
	session.Save()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "New session issued",
	})
}

func (handler *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		sessionToken := session.Get("token")
		if sessionToken == nil {
			ctx.JSON(http.StatusForbidden, gin.H{
				"message": "Not logged",
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}
