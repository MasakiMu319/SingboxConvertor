package api

import (
	"SingboxConvertor/model"
	"SingboxConvertor/service/service"
	"SingboxConvertor/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSubscription(c *gin.Context) {
	// Convertor includes all required parameters.
	Convertor := model.ConvertArg{
		Sub:       c.Request.FormValue("sub"),
		ConfigUrl: c.Request.FormValue("configurl"),
		UrlTest:   nil,
	}

	Convertor.Sub, _ = utils.Decrypt(Convertor.Sub, utils.Key)
	Convertor.ConfigUrl, _ = utils.Decrypt(Convertor.ConfigUrl, utils.Key)

	// Check subscription and external configuration are valid address.

	if !utils.IsValidURL(Convertor.Sub) || !utils.IsValidURL(Convertor.ConfigUrl) {
		c.Data(http.StatusBadRequest,
			"application/json; charset=utf-8",
			utils.GenRespJSON(http.StatusBadRequest, model.URLErr))
		return
	}

	// Check the addTag switch.
	if c.Request.FormValue("addTag") != "" {
		Convertor.AddTag = true
	}

	// Generate sing-box configuration.
	reSub, err := func() ([]byte, error) {
		Convertor.UrlTest = []model.UrlTestArg{
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
		return service.MakeConfig(c, &http.Client{}, Convertor)
	}()

	if err != nil {
		c.Data(http.StatusInternalServerError,
			"application/json; charset=utf-8",
			utils.GenRespJSON(http.StatusInternalServerError, model.GenSubErr))
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", reSub)
}
