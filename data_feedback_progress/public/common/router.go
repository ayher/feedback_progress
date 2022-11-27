package common

import (
	"data_feedback_progress/components"
	"data_feedback_progress/public/request"
	"data_feedback_progress/public/response"
	"github.com/gin-gonic/gin"
)

type RouterMap struct {
	Component func(interface{}) (interface{}, error)
	Request   request.IValidateRequest
	Response  interface{}
}

func GetRouter() map[string]RouterMap {
	return map[string]RouterMap{
		"/new_member": {
			Component: components.NewMember,
			Request:   &request.NewMemberRequest{},
			Response:  &response.NewMemberResponse{},
		},
		"/get_member": {
			Component: components.GetMember,
			Request:   &request.GetMemberRequest{},
			Response:  &response.GetMemberResponse{},
		},
		"/new_series": {
			Component: components.NewSeries,
			Request:   &request.GetSeriesRequest{},
			Response:  &response.NewSeriesResponse{},
		},
		"/get_series": {
			Component: components.GetSeries,
			Request:   &request.GetSeriesRequest{},
			Response:  &response.GetSeriesResponse{},
		},
		"/new_article": {
			Component: components.NewArticle,
			Request:   &request.NewArticleRequest{},
			Response:  &response.NewArticleResponse{},
		},
		"/get_article": {
			Component: components.GetArticle,
			Request:   &request.GetArticleRequest{},
			Response:  &response.GetArticleResponse{},
		},
	}
}

func GetComponent(component func(interface{}) (interface{}, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestLoad := request.GetRequest(ctx)
		if requestLoad == nil {
			response.False(ctx, "get req err")
			return
		}

		res, err := component(requestLoad)
		if err != nil {
			response.False(ctx, err.Error())
		} else {
			response.Sucess(ctx, res)
		}
	}

}
