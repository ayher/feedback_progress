package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FpSubscribeMgr struct {
	*_BaseMgr
}

// FpSubscribeMgr open func
func FpSubscribeMgr(db *gorm.DB) *_FpSubscribeMgr {
	if db == nil {
		panic(fmt.Errorf("FpSubscribeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FpSubscribeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fp_subscribe"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FpSubscribeMgr) GetTableName() string {
	return "fp_subscribe"
}

// Reset 重置gorm会话
func (obj *_FpSubscribeMgr) Reset() *_FpSubscribeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FpSubscribeMgr) Get() (result FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FpSubscribeMgr) Gets() (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FpSubscribeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_FpSubscribeMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithFrom from获取 订阅人
func (obj *_FpSubscribeMgr) WithFrom(from int) Option {
	return optionFunc(func(o *options) { o.query["from"] = from })
}

// WithTo to获取 被订阅人
func (obj *_FpSubscribeMgr) WithTo(to int) Option {
	return optionFunc(func(o *options) { o.query["to"] = to })
}

// WithType type获取 类型 0:用户订阅,1:系列订阅,2:文章收藏
func (obj *_FpSubscribeMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithSeriesID series_id获取 系列id
func (obj *_FpSubscribeMgr) WithSeriesID(seriesID int) Option {
	return optionFunc(func(o *options) { o.query["series_id"] = seriesID })
}

// WithArticleID article_id获取 文章id
func (obj *_FpSubscribeMgr) WithArticleID(articleID int) Option {
	return optionFunc(func(o *options) { o.query["article_id"] = articleID })
}

// WithAddTime add_time获取 记录添加时间
func (obj *_FpSubscribeMgr) WithAddTime(addTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["add_time"] = addTime })
}

// WithUpdateTime update_time获取 记录更新时间
func (obj *_FpSubscribeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_FpSubscribeMgr) GetByOption(opts ...Option) (result FpSubscribe, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FpSubscribeMgr) GetByOptions(opts ...Option) (results []*FpSubscribe, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FpSubscribeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FpSubscribe, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where(options.query)
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
func (obj *_FpSubscribeMgr) GetFromID(id int) (result FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_FpSubscribeMgr) GetBatchFromID(ids []int) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromFrom 通过from获取内容 订阅人
func (obj *_FpSubscribeMgr) GetFromFrom(from int) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`from` = ?", from).Find(&results).Error

	return
}

// GetBatchFromFrom 批量查找 订阅人
func (obj *_FpSubscribeMgr) GetBatchFromFrom(froms []int) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`from` IN (?)", froms).Find(&results).Error

	return
}

// GetFromTo 通过to获取内容 被订阅人
func (obj *_FpSubscribeMgr) GetFromTo(to int) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`to` = ?", to).Find(&results).Error

	return
}

// GetBatchFromTo 批量查找 被订阅人
func (obj *_FpSubscribeMgr) GetBatchFromTo(tos []int) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`to` IN (?)", tos).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型 0:用户订阅,1:系列订阅,2:文章收藏
func (obj *_FpSubscribeMgr) GetFromType(_type int8) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型 0:用户订阅,1:系列订阅,2:文章收藏
func (obj *_FpSubscribeMgr) GetBatchFromType(_types []int8) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromSeriesID 通过series_id获取内容 系列id
func (obj *_FpSubscribeMgr) GetFromSeriesID(seriesID int) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`series_id` = ?", seriesID).Find(&results).Error

	return
}

// GetBatchFromSeriesID 批量查找 系列id
func (obj *_FpSubscribeMgr) GetBatchFromSeriesID(seriesIDs []int) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`series_id` IN (?)", seriesIDs).Find(&results).Error

	return
}

// GetFromArticleID 通过article_id获取内容 文章id
func (obj *_FpSubscribeMgr) GetFromArticleID(articleID int) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`article_id` = ?", articleID).Find(&results).Error

	return
}

// GetBatchFromArticleID 批量查找 文章id
func (obj *_FpSubscribeMgr) GetBatchFromArticleID(articleIDs []int) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`article_id` IN (?)", articleIDs).Find(&results).Error

	return
}

// GetFromAddTime 通过add_time获取内容 记录添加时间
func (obj *_FpSubscribeMgr) GetFromAddTime(addTime time.Time) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`add_time` = ?", addTime).Find(&results).Error

	return
}

// GetBatchFromAddTime 批量查找 记录添加时间
func (obj *_FpSubscribeMgr) GetBatchFromAddTime(addTimes []time.Time) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`add_time` IN (?)", addTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 记录更新时间
func (obj *_FpSubscribeMgr) GetFromUpdateTime(updateTime time.Time) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 记录更新时间
func (obj *_FpSubscribeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FpSubscribeMgr) FetchByPrimaryKey(id int) (result FpSubscribe, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSubscribe{}).Where("`id` = ?", id).First(&result).Error

	return
}
