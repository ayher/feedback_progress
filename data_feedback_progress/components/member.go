package components

import (
	"data_feedback_progress/model"
	"data_feedback_progress/public/database"
	"data_feedback_progress/public/request"
	"data_feedback_progress/public/response"
	"fmt"
)

func NewMember(req interface{}) (interface{}, error) {

	member, ok := req.(*request.NewMemberRequest)
	if !ok {
		return nil, fmt.Errorf("parse err")
	}

	err := model.FpMemberMgr(database.GormFp()).Create(member).Error
	if err != nil {
		return nil, err
	}

	return response.NewMemberResponse{Msg: "success"}, nil
}
