package response

import "time"

type NewMemberResponse struct {
	Msg string `json:"msg"`
}

type GetMemberResponse struct {
	Name      string    `json:"name"`
	Wechat    string    `json:"wechat"`
	Avatar    string    `json:"avatar"`
	Introduce string    `json:"introduce"`
	AddTime   time.Time `json:"add_time"`
}
