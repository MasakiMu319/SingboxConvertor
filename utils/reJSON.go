package utils

import (
	"SingboxConvertor/model"
	"encoding/json"
)

// GenRespJSON generates response JSON.
func GenRespJSON(status int, errMsg string) (re []byte) {
	re, _ = json.MarshalIndent(model.Response{
		status,
		errMsg,
	}, "", "    ")
	return
}
