package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type NewArticleRequest struct {
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	SeriesId   int    `json:"series_id" binding:"required"`
	WordNumber int    `json:"word_number" binding:"required"`
}

func (r *NewArticleRequest) Validate(ctx *gin.Context) error {
	if err := ctx.ShouldBind(&r); err != nil {
		if t, ok := err.(validator.ValidationErrors); ok {
			return fmt.Errorf("invalid %s", t[0].Field())
		}
		return err
	}
	return nil
}

type GetArticleRequest struct {
	Id int `json:"id" binding:"required"`
}

func (r *GetArticleRequest) Validate(ctx *gin.Context) error {
	fmt.Println("GetArticleRequest======= 第一次有 ID，第二次没有，然后会重复")
	if err := ctx.ShouldBind(&r); err != nil {
		if t, ok := err.(validator.ValidationErrors); ok {
			return fmt.Errorf("invalid %s", t[0].Field())
		}
		return err
	}
	return nil
}

type GetRankArticleRequest struct {
}

func (r *GetRankArticleRequest) Validate(ctx *gin.Context) error {
	return nil
}
