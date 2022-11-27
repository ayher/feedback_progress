package components

import (
	"data_feedback_progress/model"
	"data_feedback_progress/public/database"
	"data_feedback_progress/public/request"
	"data_feedback_progress/public/response"
	"fmt"
)

func NewSeries(series interface{}) (interface{}, error) {
	series, ok := series.(*request.NewSeriesRequest)
	if !ok {
		return nil, fmt.Errorf("parse err")
	}

	err := model.FpSeriesMgr(database.GormFp()).Create(series).Error
	if err != nil {
		return nil, err
	}

	return response.NewSeriesResponse{Msg: "success"}, nil
}

func GetSeries(series interface{}) (interface{}, error) {
	seriesReq, ok := series.(*request.GetSeriesRequest)
	if !ok {
		return nil, fmt.Errorf("parse err")
	}

	serieDatas, err := model.FpSeriesMgr(database.GormFp()).GetFromID(seriesReq.Id)
	if err != nil {
		return nil, err
	}

	return response.GetSeriesResponse{
		SeriesName:     serieDatas.SeriesName,
		SeriesDescribe: serieDatas.SeriesDescribe,
		MemberId:       serieDatas.MemberID,
		Type:           int(serieDatas.Type),
	}, nil
}
