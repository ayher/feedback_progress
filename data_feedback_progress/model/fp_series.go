package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _FpSeriesMgr struct {
	*_BaseMgr
}

// FpSeriesMgr open func
func FpSeriesMgr(db *gorm.DB) *_FpSeriesMgr {
	if db == nil {
		panic(fmt.Errorf("FpSeriesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FpSeriesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fp_series"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FpSeriesMgr) GetTableName() string {
	return "fp_series"
}

// Reset 重置gorm会话
func (obj *_FpSeriesMgr) Reset() *_FpSeriesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FpSeriesMgr) Get() (result FpSeries, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FpSeriesMgr) Gets() (results []*FpSeries, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FpSeriesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_FpSeriesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSeriesName series_name获取 名称
func (obj *_FpSeriesMgr) WithSeriesName(seriesName string) Option {
	return optionFunc(func(o *options) { o.query["series_name"] = seriesName })
}

// WithSeriesDescribe series_describe获取 描述
func (obj *_FpSeriesMgr) WithSeriesDescribe(seriesDescribe string) Option {
	return optionFunc(func(o *options) { o.query["series_describe"] = seriesDescribe })
}

// WithMemberID member_id获取 所属人
func (obj *_FpSeriesMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithType type获取 类型 0:普通系列，1:签到
func (obj *_FpSeriesMgr) WithType(_type int8) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_FpSeriesMgr) GetByOption(opts ...Option) (result FpSeries, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FpSeriesMgr) GetByOptions(opts ...Option) (results []*FpSeries, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FpSeriesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FpSeries, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Where(options.query)
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
func (obj *_FpSeriesMgr) GetFromID(id int) (result FpSeries, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_FpSeriesMgr) GetBatchFromID(ids []int) (results []*FpSeries, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSeriesName 通过series_name获取内容 名称
func (obj *_FpSeriesMgr) GetFromSeriesName(seriesName string) (results []*FpSeries, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Where("`series_name` = ?", seriesName).Find(&results).Error

	return
}

// GetBatchFromSeriesName 批量查找 名称
func (obj *_FpSeriesMgr) GetBatchFromSeriesName(seriesNames []string) (results []*FpSeries, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Where("`series_name` IN (?)", seriesNames).Find(&results).Error

	return
}

// GetFromSeriesDescribe 通过series_describe获取内容 描述
func (obj *_FpSeriesMgr) GetFromSeriesDescribe(seriesDescribe string) (results []*FpSeries, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Where("`series_describe` = ?", seriesDescribe).Find(&results).Error

	return
}

// GetBatchFromSeriesDescribe 批量查找 描述
func (obj *_FpSeriesMgr) GetBatchFromSeriesDescribe(seriesDescribes []string) (results []*FpSeries, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Where("`series_describe` IN (?)", seriesDescribes).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 所属人
func (obj *_FpSeriesMgr) GetFromMemberID(memberID int) (results []*FpSeries, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 所属人
func (obj *_FpSeriesMgr) GetBatchFromMemberID(memberIDs []int) (results []*FpSeries, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型 0:普通系列，1:签到
func (obj *_FpSeriesMgr) GetFromType(_type int8) (results []*FpSeries, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型 0:普通系列，1:签到
func (obj *_FpSeriesMgr) GetBatchFromType(_types []int8) (results []*FpSeries, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FpSeriesMgr) FetchByPrimaryKey(id int) (result FpSeries, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpSeries{}).Where("`id` = ?", id).First(&result).Error

	return
}
