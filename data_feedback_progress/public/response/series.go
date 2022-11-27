package response

type NewSeriesResponse struct {
	Msg string `json:"msg"`
}

type GetSeriesResponse struct {
	SeriesName     string `json:"series_name"`
	SeriesDescribe string `json:"series_describe"`
	MemberId       int    `json:"member_id"`
	Type           int    `json:"type"`
}
