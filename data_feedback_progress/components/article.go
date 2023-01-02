package components

import (
	"data_feedback_progress/model"
	"data_feedback_progress/public/database"
	"data_feedback_progress/public/request"
	"data_feedback_progress/public/response"
	"fmt"
	"time"
)

func GetArticle(req interface{}) (interface{}, error) {
	articleReq, ok := req.(*request.GetArticleRequest)
	if !ok {
		return nil, fmt.Errorf("parse err")
	}

	articles, err := model.FpArticleMgr(database.GormFp()).GetFromID(articleReq.Id)
	if err != nil {
		return nil, fmt.Errorf("parse err:%+v", err)
	}

	series, errs := model.FpSeriesMgr(database.GormFp()).GetFromID(articles.SeriesID)
	if errs != nil {
		return nil, errs
	}

	member, errs := model.FpMemberMgr(database.GormFp()).GetFromID(series.MemberID)
	if errs != nil {
		return nil, errs
	}

	return response.GetArticleResponse{
		Title:      articles.Title,
		Content:    articles.Content,
		SeriesId:   series.ID,
		SeriesName: series.SeriesName,
		WordNumber: articles.WordNumber,
		MemberName: member.Name,
	}, nil
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

func GetRankArticle(req interface{}) (interface{}, error) {
	var articles []model.FpArticle
	err := model.FpArticleMgr(database.GormFp()).Order("love_number desc").Find(&articles).Error
	if err != nil {
		return nil, fmt.Errorf("parse err:%+v", err)
	}

	var res []response.GetRankArticleResponse
	for _, article := range articles {
		series, errs := model.FpSeriesMgr(database.GormFp()).GetFromID(article.SeriesID)
		if errs != nil {
			return nil, errs
		}

		member, errs := model.FpMemberMgr(database.GormFp()).GetFromID(series.MemberID)
		if errs != nil {
			return nil, errs
		}

		var subscribeCount int64
		errs = model.FpSubscribeMgr(database.GormFp()).Where("article_id = ?", article.ID).Count(&subscribeCount).Error
		if errs != nil {
			return nil, errs
		}

		var commentCount int64
		errs = model.FpCommentMgr(database.GormFp()).Where("article_id = ?", article.ID).Count(&commentCount).Error
		if errs != nil {
			return nil, errs
		}

		res = append(res, response.GetRankArticleResponse{
			Id:         article.ID,
			Title:      article.Title,
			MemberName: member.Name,
			Love:       article.LoveNumber,
			Collect:    int(subscribeCount),
			Comment:    int(commentCount),
		})
	}

	return res, nil
}
