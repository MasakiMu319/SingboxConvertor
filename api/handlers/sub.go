package handlers

import (
	model2 "SingboxConvertor/api/model"
	"SingboxConvertor/service/service"
	"SingboxConvertor/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SubHandler struct{}

func (handler *SubHandler) GenSubHandler(ctx *gin.Context) {
	var subData model2.SubArg

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

func (handler *SubHandler) GetSubHandler(ctx *gin.Context) {
	// Convertor includes all required parameters.
	Convertor := model2.ConvertArg{
		Sub:       ctx.Request.FormValue("sub"),
		ConfigUrl: ctx.Request.FormValue("configurl"),
		UrlTest:   nil,
	}

	Convertor.Sub, _ = utils.Decrypt(Convertor.Sub, utils.Key)
	Convertor.ConfigUrl, _ = utils.Decrypt(Convertor.ConfigUrl, utils.Key)

	// Check subscription and external configuration are valid address.

	if !utils.IsValidURL(Convertor.Sub) || !utils.IsValidURL(Convertor.ConfigUrl) {
		ctx.Data(http.StatusBadRequest,
			"application/json; charset=utf-8",
			utils.GenRespJSON(http.StatusBadRequest, model2.URLErr))
		return
	}

	// Check the addTag switch.
	if ctx.Request.FormValue("addTag") != "" {
		Convertor.AddTag = true
	}

	// Generate sing-box configuration.
	reSub, err := func() ([]byte, error) {
		Convertor.UrlTest = []model2.UrlTestArg{
			{
				Tag:     "HK",
				Include: "HK|HongKong|🇭🇰|香港",
				Type:    "selector",
			},
			{
				Tag:     "TW",
				Include: "TW|Taiwan|🇹🇼|台湾|台灣",
				Type:    "selector",
			},
			{
				Tag:     "JP",
				Include: "JP|Japan|🇯🇵|日本",
				Type:    "selector",
			},
			{
				Tag:     "SG",
				Include: "SG|Singapore|🇸🇬|新加坡|獅城",
				Type:    "selector",
			},
			{
				Tag:     "US",
				Include: "US|United States|🇺🇸|美国|美國",
				Type:    "selector",
			},
			{
				Tag:  "fallback",
				Type: "selector",
			},
		}
		return service.MakeConfig(ctx, &http.Client{}, Convertor)
	}()

	if err != nil {
		ctx.Data(http.StatusInternalServerError,
			"application/json; charset=utf-8",
			utils.GenRespJSON(http.StatusInternalServerError, model2.GenSubErr))
		return
	}

	ctx.Data(http.StatusOK, "application/json; charset=utf-8", reSub)
}
