package components

import (
	"data_feedback_progress/model"
	"data_feedback_progress/public/database"
	"data_feedback_progress/public/request"
	"data_feedback_progress/public/response"
	"fmt"
	"time"
)

func GetArticle(interface{}) (interface{}, error) {
	return "success", nil
}

func NewArticle(req interface{}) (interface{}, error) {
	article, ok := req.(*request.NewArticleRequest)
	if !ok {
		return nil, fmt.Errorf("parse err")
	}

	var now = time.Now()

	err := model.FpArticleMgr(database.GormFp()).Create(struct {
		Title       string    `json:"title" binding:"required"`
		Content     string    `json:"content" binding:"required"`
		SeriesId    int       `json:"series_id" binding:"required"`
		WordNumber  int       `json:"word_number" binding:"required"`
		PublishTime time.Time `json:"publish_time"`
	}{
		Title:       article.Title,
		Content:     article.Content,
		SeriesId:    article.SeriesId,
		WordNumber:  article.WordNumber,
		PublishTime: now,
	}).Error
	if err != nil {
		return nil, err
	}

	return response.NewArticleResponse{Msg: "success"}, nil
}
