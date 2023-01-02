package response

type NewArticleResponse struct {
	Msg string `json:"msg"`
}

type GetArticleResponse struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	SeriesId   int    `json:"series_id"`
	SeriesName   string    `json:"series_name"`
	WordNumber int    `json:"word_number"`
	MemberName string `json:"member_name"`
}

type GetRankArticleResponse struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	MemberName string `json:"member_name"`
	Love       int    `json:"love"`
	Collect    int    `json:"collect"`
	Comment    int    `json:"comment"`
}
