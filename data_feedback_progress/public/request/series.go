package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type NewSeriesRequest struct {
	SeriesName     string `json:"series_name" binding:"required"`
	SeriesDescribe string `json:"series_describe" binding:"required"`
	MemberId       int    `json:"member_id" binding:"required"`
	Type           int    `json:"type" binding:"required"`
}

func (r *NewSeriesRequest) Validate(ctx *gin.Context) error {
	if err := ctx.ShouldBind(&r); err != nil {
		if t, ok := err.(validator.ValidationErrors); ok {
			return fmt.Errorf("invalid %s", t[0].Field())
		}
		return err
	}
	return nil
}
