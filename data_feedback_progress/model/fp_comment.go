package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FpCommentMgr struct {
	*_BaseMgr
}

// FpCommentMgr open func
func FpCommentMgr(db *gorm.DB) *_FpCommentMgr {
	if db == nil {
		panic(fmt.Errorf("FpCommentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FpCommentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fp_comment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FpCommentMgr) GetTableName() string {
	return "fp_comment"
}

// Reset 重置gorm会话
func (obj *_FpCommentMgr) Reset() *_FpCommentMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FpCommentMgr) Get() (result FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FpCommentMgr) Gets() (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FpCommentMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FpComment{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_FpCommentMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithArticleID article_id获取 文章 id
func (obj *_FpCommentMgr) WithArticleID(articleID int) Option {
	return optionFunc(func(o *options) { o.query["article_id"] = articleID })
}

// WithContext context获取 内容
func (obj *_FpCommentMgr) WithContext(context string) Option {
	return optionFunc(func(o *options) { o.query["context"] = context })
}

// WithMemberID member_id获取 评论人
func (obj *_FpCommentMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithCommentID comment_id获取 父评论 id -1:无父评论
func (obj *_FpCommentMgr) WithCommentID(commentID int) Option {
	return optionFunc(func(o *options) { o.query["comment_id"] = commentID })
}

// WithLove love获取 点赞数
func (obj *_FpCommentMgr) WithLove(love int) Option {
	return optionFunc(func(o *options) { o.query["love"] = love })
}

// WithAddTime add_time获取 记录添加时间
func (obj *_FpCommentMgr) WithAddTime(addTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["add_time"] = addTime })
}

// WithUpdateTime update_time获取 记录更新时间
func (obj *_FpCommentMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_FpCommentMgr) GetByOption(opts ...Option) (result FpComment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FpCommentMgr) GetByOptions(opts ...Option) (results []*FpComment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FpCommentMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FpComment, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_FpCommentMgr) GetFromID(id int) (result FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_FpCommentMgr) GetBatchFromID(ids []int) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromArticleID 通过article_id获取内容 文章 id
func (obj *_FpCommentMgr) GetFromArticleID(articleID int) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`article_id` = ?", articleID).Find(&results).Error

	return
}

// GetBatchFromArticleID 批量查找 文章 id
func (obj *_FpCommentMgr) GetBatchFromArticleID(articleIDs []int) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`article_id` IN (?)", articleIDs).Find(&results).Error

	return
}

// GetFromContext 通过context获取内容 内容
func (obj *_FpCommentMgr) GetFromContext(context string) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`context` = ?", context).Find(&results).Error

	return
}

// GetBatchFromContext 批量查找 内容
func (obj *_FpCommentMgr) GetBatchFromContext(contexts []string) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`context` IN (?)", contexts).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 评论人
func (obj *_FpCommentMgr) GetFromMemberID(memberID int) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 评论人
func (obj *_FpCommentMgr) GetBatchFromMemberID(memberIDs []int) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromCommentID 通过comment_id获取内容 父评论 id -1:无父评论
func (obj *_FpCommentMgr) GetFromCommentID(commentID int) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`comment_id` = ?", commentID).Find(&results).Error

	return
}

// GetBatchFromCommentID 批量查找 父评论 id -1:无父评论
func (obj *_FpCommentMgr) GetBatchFromCommentID(commentIDs []int) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`comment_id` IN (?)", commentIDs).Find(&results).Error

	return
}

// GetFromLove 通过love获取内容 点赞数
func (obj *_FpCommentMgr) GetFromLove(love int) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`love` = ?", love).Find(&results).Error

	return
}

// GetBatchFromLove 批量查找 点赞数
func (obj *_FpCommentMgr) GetBatchFromLove(loves []int) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`love` IN (?)", loves).Find(&results).Error

	return
}

// GetFromAddTime 通过add_time获取内容 记录添加时间
func (obj *_FpCommentMgr) GetFromAddTime(addTime time.Time) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`add_time` = ?", addTime).Find(&results).Error

	return
}

// GetBatchFromAddTime 批量查找 记录添加时间
func (obj *_FpCommentMgr) GetBatchFromAddTime(addTimes []time.Time) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`add_time` IN (?)", addTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 记录更新时间
func (obj *_FpCommentMgr) GetFromUpdateTime(updateTime time.Time) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 记录更新时间
func (obj *_FpCommentMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FpCommentMgr) FetchByPrimaryKey(id int) (result FpComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpComment{}).Where("`id` = ?", id).First(&result).Error

	return
}
