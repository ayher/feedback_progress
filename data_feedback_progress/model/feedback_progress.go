package model

import (
	"gorm.io/datatypes"
	"time"
)

// FpArticle 文章
type FpArticle struct {
	ID          int       `gorm:"primaryKey;column:id" json:"-"`
	Title       string    `gorm:"column:title" json:"title"`              // 标题
	Content     string    `gorm:"column:content" json:"content"`          // 内容
	SeriesID    int       `gorm:"column:series_id" json:"seriesId"`       // 所属系列
	LoveNumber  int       `gorm:"column:love_number" json:"loveNumber"`   // 点赞数
	WordNumber  int       `gorm:"column:word_number" json:"wordNumber"`   // 字数
	PublishTime time.Time `gorm:"column:publish_time" json:"publishTime"` // 发布时间
	AddTime     time.Time `gorm:"column:add_time" json:"addTime"`         // 记录添加时间
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`   // 记录更新时间
}

// TableName get sql table name.获取数据库表名
func (m *FpArticle) TableName() string {
	return "fp_article"
}

// FpArticleColumns get sql column name.获取数据库列名
var FpArticleColumns = struct {
	ID          string
	Title       string
	Content     string
	SeriesID    string
	LoveNumber  string
	WordNumber  string
	PublishTime string
	AddTime     string
	UpdateTime  string
}{
	ID:          "id",
	Title:       "title",
	Content:     "content",
	SeriesID:    "series_id",
	LoveNumber:  "love_number",
	WordNumber:  "word_number",
	PublishTime: "publish_time",
	AddTime:     "add_time",
	UpdateTime:  "update_time",
}

// FpClockIn 签到表
type FpClockIn struct {
	ID         int            `gorm:"primaryKey;column:id" json:"-"`
	ArticleID  int            `gorm:"column:article_id" json:"articleId"`   // 文章 id
	ClockDate  datatypes.Date `gorm:"column:clock_date" json:"clockDate"`   // 签到对象日期
	Score      int            `gorm:"column:score" json:"score"`            // 评分
	AddTime    time.Time      `gorm:"column:add_time" json:"addTime"`       // 记录添加时间
	UpdateTime time.Time      `gorm:"column:update_time" json:"updateTime"` // 记录更新时间
}

// TableName get sql table name.获取数据库表名
func (m *FpClockIn) TableName() string {
	return "fp_clock_in"
}

// FpClockInColumns get sql column name.获取数据库列名
var FpClockInColumns = struct {
	ID         string
	ArticleID  string
	ClockDate  string
	Score      string
	AddTime    string
	UpdateTime string
}{
	ID:         "id",
	ArticleID:  "article_id",
	ClockDate:  "clock_date",
	Score:      "score",
	AddTime:    "add_time",
	UpdateTime: "update_time",
}

// FpComment 评论表
type FpComment struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`
	ArticleID  int       `gorm:"column:article_id" json:"articleId"`   // 文章 id
	Context    string    `gorm:"column:context" json:"context"`        // 内容
	MemberID   int       `gorm:"column:member_id" json:"memberId"`     // 评论人
	CommentID  int       `gorm:"column:comment_id" json:"commentId"`   // 父评论 id -1:无父评论
	Love       int       `gorm:"column:love" json:"love"`              // 点赞数
	AddTime    time.Time `gorm:"column:add_time" json:"addTime"`       // 记录添加时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 记录更新时间
}

// TableName get sql table name.获取数据库表名
func (m *FpComment) TableName() string {
	return "fp_comment"
}

// FpCommentColumns get sql column name.获取数据库列名
var FpCommentColumns = struct {
	ID         string
	ArticleID  string
	Context    string
	MemberID   string
	CommentID  string
	Love       string
	AddTime    string
	UpdateTime string
}{
	ID:         "id",
	ArticleID:  "article_id",
	Context:    "context",
	MemberID:   "member_id",
	CommentID:  "comment_id",
	Love:       "love",
	AddTime:    "add_time",
	UpdateTime: "update_time",
}

// FpMember 用户表
type FpMember struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`
	Name       string    `gorm:"column:name" json:"name"`              // 昵称
	Wechat     string    `gorm:"column:wechat" json:"wechat"`          // 微信
	Avatar     string    `gorm:"column:avatar" json:"avatar"`          // 头像
	Introduce  string    `gorm:"column:introduce" json:"introduce"`    // 介绍
	AddTime    time.Time `gorm:"column:add_time" json:"addTime"`       // 记录添加时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 记录更新时间
}

// TableName get sql table name.获取数据库表名
func (m *FpMember) TableName() string {
	return "fp_member"
}

// FpMemberColumns get sql column name.获取数据库列名
var FpMemberColumns = struct {
	ID         string
	Name       string
	Wechat     string
	Avatar     string
	Introduce  string
	AddTime    string
	UpdateTime string
}{
	ID:         "id",
	Name:       "name",
	Wechat:     "wechat",
	Avatar:     "avatar",
	Introduce:  "introduce",
	AddTime:    "add_time",
	UpdateTime: "update_time",
}

// FpSeries 系列
type FpSeries struct {
	ID             int    `gorm:"primaryKey;column:id" json:"-"`
	SeriesName     string `gorm:"column:series_name" json:"seriesName"`         // 名称
	SeriesDescribe string `gorm:"column:series_describe" json:"seriesDescribe"` // 描述
	MemberID       int    `gorm:"column:member_id" json:"memberId"`             // 所属人
	Type           int8   `gorm:"column:type" json:"type"`                      // 类型 0:普通系列，1:签到
}

// TableName get sql table name.获取数据库表名
func (m *FpSeries) TableName() string {
	return "fp_series"
}

// FpSeriesColumns get sql column name.获取数据库列名
var FpSeriesColumns = struct {
	ID             string
	SeriesName     string
	SeriesDescribe string
	MemberID       string
	Type           string
}{
	ID:             "id",
	SeriesName:     "series_name",
	SeriesDescribe: "series_describe",
	MemberID:       "member_id",
	Type:           "type",
}

// FpSubscribe 订阅表
type FpSubscribe struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`
	From       int       `gorm:"column:from" json:"from"`              // 订阅人
	To         int       `gorm:"column:to" json:"to"`                  // 被订阅人
	Type       int8      `gorm:"column:type" json:"type"`              // 类型 0:用户订阅,1:系列订阅,2:文章收藏
	SeriesID   int       `gorm:"column:series_id" json:"seriesId"`     // 系列id
	ArticleID  int       `gorm:"column:article_id" json:"articleId"`   // 文章id
	AddTime    time.Time `gorm:"column:add_time" json:"addTime"`       // 记录添加时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 记录更新时间
}

// TableName get sql table name.获取数据库表名
func (m *FpSubscribe) TableName() string {
	return "fp_subscribe"
}

// FpSubscribeColumns get sql column name.获取数据库列名
var FpSubscribeColumns = struct {
	ID         string
	From       string
	To         string
	Type       string
	SeriesID   string
	ArticleID  string
	AddTime    string
	UpdateTime string
}{
	ID:         "id",
	From:       "from",
	To:         "to",
	Type:       "type",
	SeriesID:   "series_id",
	ArticleID:  "article_id",
	AddTime:    "add_time",
	UpdateTime: "update_time",
}

// FpTarget 目标表
type FpTarget struct {
	ID               int       `gorm:"primaryKey;column:id" json:"-"`
	ClockInID        int       `gorm:"column:clock_in_id" json:"clockInId"`              // 签到 id
	Context          string    `gorm:"column:context" json:"context"`                    // 目标内容
	Duration         int       `gorm:"column:duration" json:"duration"`                  // 耗时分钟
	ConcentrateScore int8      `gorm:"column:concentrate_score" json:"concentrateScore"` // 专注评分
	HarvestScore     int8      `gorm:"column:harvest_score" json:"harvestScore"`         // 收获评分
	AddTime          time.Time `gorm:"column:add_time" json:"addTime"`                   // 记录添加时间
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`             // 记录更新时间
}

// TableName get sql table name.获取数据库表名
func (m *FpTarget) TableName() string {
	return "fp_target"
}

// FpTargetColumns get sql column name.获取数据库列名
var FpTargetColumns = struct {
	ID               string
	ClockInID        string
	Context          string
	Duration         string
	ConcentrateScore string
	HarvestScore     string
	AddTime          string
	UpdateTime       string
}{
	ID:               "id",
	ClockInID:        "clock_in_id",
	Context:          "context",
	Duration:         "duration",
	ConcentrateScore: "concentrate_score",
	HarvestScore:     "harvest_score",
	AddTime:          "add_time",
	UpdateTime:       "update_time",
}
