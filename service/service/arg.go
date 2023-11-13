package service

import (
	"SingboxConvertor/model"
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xmdhs/clash2singbox/httputils"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"SingboxConvertor/db"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"lukechampine.com/blake3"
)

func PutArg(cxt context.Context, arg model.ConvertArg, db db.DB) (string, error) {
	b, err := json.Marshal(arg)
	if err != nil {
		return "", fmt.Errorf("PutArg: %w", err)
	}
	hash := blake3.Sum256(b)
	h := hex.EncodeToString(hash[:])
	err = db.PutArg(cxt, h, arg)
	if err != nil {
		return "", fmt.Errorf("PutArg: %w", err)
	}
	return h, nil
}

func GetSub(cxt *gin.Context, c *http.Client, db db.DB, id string) ([]byte, error) {
	arg, err := db.GetArg(cxt, id)
	if err != nil {
		return nil, fmt.Errorf("GetSub: %w", err)
	}
	b, err := MakeConfig(cxt, c, arg)
	if err != nil {
		return nil, fmt.Errorf("GetSub: %w", err)
	}
	return b, nil
}

func MakeConfig(cxt *gin.Context, c *http.Client, arg model.ConvertArg) ([]byte, error) {
	//Get configuration.
	if arg.ConfigUrl != "" {
		b, err := httputils.HttpGet(cxt, c, arg.ConfigUrl, 1000*1000*10)
		log.Println("External configuration url:", arg.ConfigUrl)
		if err != nil {
			log.Println("Get configuration error:", err)
			return nil, fmt.Errorf("MakeConfig: %w", err)
		}
		log.Println("Get configuration succeed")
		arg.Config = string(b)
	}

	// Convert.
	b, err := convert2sing(cxt, c, arg.Config, arg.Sub, arg.Include, arg.Exclude, arg.AddTag)
	if err != nil {
		log.Println("Convert failed:", err)
		return nil, fmt.Errorf("MakeConfig: %w", err)
	}
	log.Println("The Sing-box configuration has been generated succeed")

	// Add custom groups.
	if len(arg.UrlTest) != 0 {
		log.Println("Adding custom groups...")
		nb, err := customUrlTest(b, arg.UrlTest)
		if err != nil {
			log.Println("Add custom groups error.")
			return nil, fmt.Errorf("MakeConfig: %w", err)
		}
		b = nb
		log.Println("Add custom groups succeed.")
	}
	log.Println("============== Separator ==============")
	return b, nil
}

var (
	ErrJson = errors.New("wrong json")
)

func customUrlTest(config []byte, u []model.UrlTestArg) ([]byte, error) {
	r := gjson.GetBytes(config, `outbounds.#(tag=="urltest").outbounds`)
	if !r.Exists() {
		return nil, fmt.Errorf("customUrlTest: %w", ErrJson)
	}
	sl := []model.SingUrltest{}

	tags := []string{}
	r.ForEach(func(key, value gjson.Result) bool {
		tags = append(tags, value.String())
		return true
	})

	for _, v := range u {
		nt, err := filter(v.Include, tags, true)
		if err != nil {
			return nil, fmt.Errorf("customUrlTest: %w", err)
		}
		nt, err = filter(v.Exclude, nt, false)
		if err != nil {
			return nil, fmt.Errorf("customUrlTest: %w", err)
		}
		var t int
		if v.Type == "urltest" {
			t, _ = lo.TryOr[int](func() (int, error) { return strconv.Atoi(v.Tolerance) }, 0)
		}
		if v.Type == "" {
			v.Type = "urltest"
		}
		sl = append(sl, model.SingUrltest{
			Outbounds: nt,
			Tag:       v.Tag,
			Tolerance: t,
			Type:      v.Type,
		})
	}

	for _, v := range sl {
		var err error
		v := v
		config, err = sjson.SetBytes(config, "outbounds.-1", v)
		if err != nil {
			return nil, fmt.Errorf("customUrlTest: %w", err)
		}
	}
	var a any
	lo.Must0(json.Unmarshal(config, &a))
	bw := bytes.NewBuffer(nil)
	jw := json.NewEncoder(bw)
	jw.SetEscapeHTML(false)
	jw.SetIndent("", "    ")
	lo.Must0(jw.Encode(a))
	return bw.Bytes(), nil
}

func filter(reg string, tags []string, need bool) ([]string, error) {
	if reg == "" {
		return tags, nil
	}
	r, err := regexp.Compile(reg)
	if err != nil {
		return nil, fmt.Errorf("filter: %w", err)
	}
	tag := lo.Filter[string](tags, func(item string, index int) bool {
		has := r.MatchString(item)
		return has == need
	})
	return tag, nil
}
