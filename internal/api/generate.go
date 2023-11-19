package api

import (
	"SingboxConvertor/model"
	"SingboxConvertor/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostGenerate(ctx *gin.Context) {
	var subData model.SubArg

	if err := ctx.ShouldBindJSON(&subData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	sub, _ := utils.Encrypt(subData.Sub, utils.Key)
	configUrl, _ := utils.Encrypt(subData.ConfigUrl, utils.Key)

	scheme := "http" // 默认使用http，如果您的服务支持https，请相应地更改
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	host := ctx.Request.Host

	url := fmt.Sprint(scheme, "://", host, "/sub?configurl=", configUrl, "&sub=", sub)

	// TODO: Check later.
	if subData.AddTag {
		url = fmt.Sprint(url, "&addTag=true")
	}

	fmt.Println(url)

	ctx.JSON(http.StatusOK, gin.H{
		"result": url,
	})
}
