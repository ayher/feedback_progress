package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FpArticleMgr struct {
	*_BaseMgr
}

// FpArticleMgr open func
func FpArticleMgr(db *gorm.DB) *_FpArticleMgr {
	if db == nil {
		panic(fmt.Errorf("FpArticleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FpArticleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fp_article"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FpArticleMgr) GetTableName() string {
	return "fp_article"
}

// Reset 重置gorm会话
func (obj *_FpArticleMgr) Reset() *_FpArticleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FpArticleMgr) Get() (result FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FpArticleMgr) Gets() (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FpArticleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_FpArticleMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取 标题
func (obj *_FpArticleMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithContent content获取 内容
func (obj *_FpArticleMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithSeriesID series_id获取 所属系列
func (obj *_FpArticleMgr) WithSeriesID(seriesID int) Option {
	return optionFunc(func(o *options) { o.query["series_id"] = seriesID })
}

// WithLoveNumber love_number获取 点赞数
func (obj *_FpArticleMgr) WithLoveNumber(loveNumber int) Option {
	return optionFunc(func(o *options) { o.query["love_number"] = loveNumber })
}

// WithWordNumber word_number获取 字数
func (obj *_FpArticleMgr) WithWordNumber(wordNumber int) Option {
	return optionFunc(func(o *options) { o.query["word_number"] = wordNumber })
}

// WithPublishTime publish_time获取 发布时间
func (obj *_FpArticleMgr) WithPublishTime(publishTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["publish_time"] = publishTime })
}

// WithAddTime add_time获取 记录添加时间
func (obj *_FpArticleMgr) WithAddTime(addTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["add_time"] = addTime })
}

// WithUpdateTime update_time获取 记录更新时间
func (obj *_FpArticleMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_FpArticleMgr) GetByOption(opts ...Option) (result FpArticle, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FpArticleMgr) GetByOptions(opts ...Option) (results []*FpArticle, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FpArticleMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FpArticle, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where(options.query)
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
func (obj *_FpArticleMgr) GetFromID(id int) (result FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_FpArticleMgr) GetBatchFromID(ids []int) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_FpArticleMgr) GetFromTitle(title string) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *_FpArticleMgr) GetBatchFromTitle(titles []string) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 内容
func (obj *_FpArticleMgr) GetFromContent(content string) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 内容
func (obj *_FpArticleMgr) GetBatchFromContent(contents []string) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromSeriesID 通过series_id获取内容 所属系列
func (obj *_FpArticleMgr) GetFromSeriesID(seriesID int) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`series_id` = ?", seriesID).Find(&results).Error

	return
}

// GetBatchFromSeriesID 批量查找 所属系列
func (obj *_FpArticleMgr) GetBatchFromSeriesID(seriesIDs []int) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`series_id` IN (?)", seriesIDs).Find(&results).Error

	return
}

// GetFromLoveNumber 通过love_number获取内容 点赞数
func (obj *_FpArticleMgr) GetFromLoveNumber(loveNumber int) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`love_number` = ?", loveNumber).Find(&results).Error

	return
}

// GetBatchFromLoveNumber 批量查找 点赞数
func (obj *_FpArticleMgr) GetBatchFromLoveNumber(loveNumbers []int) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`love_number` IN (?)", loveNumbers).Find(&results).Error

	return
}

// GetFromWordNumber 通过word_number获取内容 字数
func (obj *_FpArticleMgr) GetFromWordNumber(wordNumber int) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`word_number` = ?", wordNumber).Find(&results).Error

	return
}

// GetBatchFromWordNumber 批量查找 字数
func (obj *_FpArticleMgr) GetBatchFromWordNumber(wordNumbers []int) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`word_number` IN (?)", wordNumbers).Find(&results).Error

	return
}

// GetFromPublishTime 通过publish_time获取内容 发布时间
func (obj *_FpArticleMgr) GetFromPublishTime(publishTime time.Time) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`publish_time` = ?", publishTime).Find(&results).Error

	return
}

// GetBatchFromPublishTime 批量查找 发布时间
func (obj *_FpArticleMgr) GetBatchFromPublishTime(publishTimes []time.Time) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`publish_time` IN (?)", publishTimes).Find(&results).Error

	return
}

// GetFromAddTime 通过add_time获取内容 记录添加时间
func (obj *_FpArticleMgr) GetFromAddTime(addTime time.Time) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`add_time` = ?", addTime).Find(&results).Error

	return
}

// GetBatchFromAddTime 批量查找 记录添加时间
func (obj *_FpArticleMgr) GetBatchFromAddTime(addTimes []time.Time) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`add_time` IN (?)", addTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 记录更新时间
func (obj *_FpArticleMgr) GetFromUpdateTime(updateTime time.Time) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 记录更新时间
func (obj *_FpArticleMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FpArticleMgr) FetchByPrimaryKey(id int) (result FpArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpArticle{}).Where("`id` = ?", id).First(&result).Error

	return
}
