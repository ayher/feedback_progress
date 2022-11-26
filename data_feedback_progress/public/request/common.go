package request

import (
	"github.com/gin-gonic/gin"
)

const (
	RequestCtxKey = "request"
)

type IRequest interface{}

type IValidateRequest interface {
	IRequest
	Validate(ctx *gin.Context) error
}

func SetRequest(ctx *gin.Context, iReq IValidateRequest) {
	ctx.Set(RequestCtxKey, iReq)
}

func GetRequest(ctx *gin.Context) IRequest {
	req, exist := ctx.Get(RequestCtxKey)
	if !exist {
		return nil
	}

	iReq, ok := req.(IRequest)
	if !ok {
		return nil
	}

	return iReq
}
