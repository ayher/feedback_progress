package response

type NewArticleResponse struct {
	Msg string `json:"msg"`
}

type GetArticleResponse struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	SeriesId   int    `json:"series_id"`
	WordNumber int    `json:"word_number"`
}
