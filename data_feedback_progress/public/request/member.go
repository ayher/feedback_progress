package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type NewMemberRequest struct {
	Name      string `json:"name" binding:"required"`
	Wechat    string `json:"wechat" binding:"required"`
	Avatar    string `json:"avatar" binding:"required"`
	Introduce string `json:"introduce" binding:"required"`
}

func (r *NewMemberRequest) Validate(ctx *gin.Context) error {
	if err := ctx.ShouldBind(&r); err != nil {
		if t, ok := err.(validator.ValidationErrors); ok {
			return fmt.Errorf("invalid %s", t[0].Field())
		}
		return err
	}
	return nil
}
