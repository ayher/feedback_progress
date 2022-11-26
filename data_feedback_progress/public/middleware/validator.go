package middleware

import (
	"data_feedback_progress/public/common"
	"data_feedback_progress/public/request"
	"data_feedback_progress/public/response"
	"github.com/gin-gonic/gin"
)

func Validate(routerMap map[string]common.RouterMap) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		v, exist := routerMap[ctx.Request.URL.Path]
		if !exist {
			return
		}

		iReq := v.Request
		err := iReq.Validate(ctx)
		if err != nil {
			response.False(ctx, err.Error())
		} else {
			request.SetRequest(ctx, iReq)
		}
	}
}
