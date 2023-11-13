package db

import (
	"context"

	"SingboxConvertor/model"
)

type DB interface {
	GetArg(cxt context.Context, blake3 string) (model.ConvertArg, error)
	PutArg(cxt context.Context, blake3 string, arg model.ConvertArg) error
}
