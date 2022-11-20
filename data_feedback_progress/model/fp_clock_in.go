package model

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type _FpClockInMgr struct {
	*_BaseMgr
}

// FpClockInMgr open func
func FpClockInMgr(db *gorm.DB) *_FpClockInMgr {
	if db == nil {
		panic(fmt.Errorf("FpClockInMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FpClockInMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fp_clock_in"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FpClockInMgr) GetTableName() string {
	return "fp_clock_in"
}

// Reset 重置gorm会话
func (obj *_FpClockInMgr) Reset() *_FpClockInMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FpClockInMgr) Get() (result FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FpClockInMgr) Gets() (results []*FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FpClockInMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_FpClockInMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithArticleID article_id获取 文章 id
func (obj *_FpClockInMgr) WithArticleID(articleID int) Option {
	return optionFunc(func(o *options) { o.query["article_id"] = articleID })
}

// WithClockDate clock_date获取 签到对象日期
func (obj *_FpClockInMgr) WithClockDate(clockDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["clock_date"] = clockDate })
}

// WithScore score获取 评分
func (obj *_FpClockInMgr) WithScore(score int) Option {
	return optionFunc(func(o *options) { o.query["score"] = score })
}

// WithAddTime add_time获取 记录添加时间
func (obj *_FpClockInMgr) WithAddTime(addTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["add_time"] = addTime })
}

// WithUpdateTime update_time获取 记录更新时间
func (obj *_FpClockInMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_FpClockInMgr) GetByOption(opts ...Option) (result FpClockIn, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FpClockInMgr) GetByOptions(opts ...Option) (results []*FpClockIn, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FpClockInMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FpClockIn, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where(options.query)
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
func (obj *_FpClockInMgr) GetFromID(id int) (result FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_FpClockInMgr) GetBatchFromID(ids []int) (results []*FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromArticleID 通过article_id获取内容 文章 id
func (obj *_FpClockInMgr) GetFromArticleID(articleID int) (results []*FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where("`article_id` = ?", articleID).Find(&results).Error

	return
}

// GetBatchFromArticleID 批量查找 文章 id
func (obj *_FpClockInMgr) GetBatchFromArticleID(articleIDs []int) (results []*FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where("`article_id` IN (?)", articleIDs).Find(&results).Error

	return
}

// GetFromClockDate 通过clock_date获取内容 签到对象日期
func (obj *_FpClockInMgr) GetFromClockDate(clockDate datatypes.Date) (results []*FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where("`clock_date` = ?", clockDate).Find(&results).Error

	return
}

// GetBatchFromClockDate 批量查找 签到对象日期
func (obj *_FpClockInMgr) GetBatchFromClockDate(clockDates []datatypes.Date) (results []*FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where("`clock_date` IN (?)", clockDates).Find(&results).Error

	return
}

// GetFromScore 通过score获取内容 评分
func (obj *_FpClockInMgr) GetFromScore(score int) (results []*FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where("`score` = ?", score).Find(&results).Error

	return
}

// GetBatchFromScore 批量查找 评分
func (obj *_FpClockInMgr) GetBatchFromScore(scores []int) (results []*FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where("`score` IN (?)", scores).Find(&results).Error

	return
}

// GetFromAddTime 通过add_time获取内容 记录添加时间
func (obj *_FpClockInMgr) GetFromAddTime(addTime time.Time) (results []*FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where("`add_time` = ?", addTime).Find(&results).Error

	return
}

// GetBatchFromAddTime 批量查找 记录添加时间
func (obj *_FpClockInMgr) GetBatchFromAddTime(addTimes []time.Time) (results []*FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where("`add_time` IN (?)", addTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 记录更新时间
func (obj *_FpClockInMgr) GetFromUpdateTime(updateTime time.Time) (results []*FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 记录更新时间
func (obj *_FpClockInMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FpClockInMgr) FetchByPrimaryKey(id int) (result FpClockIn, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpClockIn{}).Where("`id` = ?", id).First(&result).Error

	return
}
