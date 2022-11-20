package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FpTargetMgr struct {
	*_BaseMgr
}

// FpTargetMgr open func
func FpTargetMgr(db *gorm.DB) *_FpTargetMgr {
	if db == nil {
		panic(fmt.Errorf("FpTargetMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FpTargetMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fp_target"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FpTargetMgr) GetTableName() string {
	return "fp_target"
}

// Reset 重置gorm会话
func (obj *_FpTargetMgr) Reset() *_FpTargetMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FpTargetMgr) Get() (result FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FpTargetMgr) Gets() (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FpTargetMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_FpTargetMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithClockInID clock_in_id获取 签到 id
func (obj *_FpTargetMgr) WithClockInID(clockInID int) Option {
	return optionFunc(func(o *options) { o.query["clock_in_id"] = clockInID })
}

// WithContext context获取 目标内容
func (obj *_FpTargetMgr) WithContext(context string) Option {
	return optionFunc(func(o *options) { o.query["context"] = context })
}

// WithDuration duration获取 耗时分钟
func (obj *_FpTargetMgr) WithDuration(duration int) Option {
	return optionFunc(func(o *options) { o.query["duration"] = duration })
}

// WithConcentrateScore concentrate_score获取 专注评分
func (obj *_FpTargetMgr) WithConcentrateScore(concentrateScore int8) Option {
	return optionFunc(func(o *options) { o.query["concentrate_score"] = concentrateScore })
}

// WithHarvestScore harvest_score获取 收获评分
func (obj *_FpTargetMgr) WithHarvestScore(harvestScore int8) Option {
	return optionFunc(func(o *options) { o.query["harvest_score"] = harvestScore })
}

// WithAddTime add_time获取 记录添加时间
func (obj *_FpTargetMgr) WithAddTime(addTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["add_time"] = addTime })
}

// WithUpdateTime update_time获取 记录更新时间
func (obj *_FpTargetMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_FpTargetMgr) GetByOption(opts ...Option) (result FpTarget, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FpTargetMgr) GetByOptions(opts ...Option) (results []*FpTarget, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FpTargetMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FpTarget, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where(options.query)
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
func (obj *_FpTargetMgr) GetFromID(id int) (result FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_FpTargetMgr) GetBatchFromID(ids []int) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromClockInID 通过clock_in_id获取内容 签到 id
func (obj *_FpTargetMgr) GetFromClockInID(clockInID int) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`clock_in_id` = ?", clockInID).Find(&results).Error

	return
}

// GetBatchFromClockInID 批量查找 签到 id
func (obj *_FpTargetMgr) GetBatchFromClockInID(clockInIDs []int) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`clock_in_id` IN (?)", clockInIDs).Find(&results).Error

	return
}

// GetFromContext 通过context获取内容 目标内容
func (obj *_FpTargetMgr) GetFromContext(context string) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`context` = ?", context).Find(&results).Error

	return
}

// GetBatchFromContext 批量查找 目标内容
func (obj *_FpTargetMgr) GetBatchFromContext(contexts []string) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`context` IN (?)", contexts).Find(&results).Error

	return
}

// GetFromDuration 通过duration获取内容 耗时分钟
func (obj *_FpTargetMgr) GetFromDuration(duration int) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`duration` = ?", duration).Find(&results).Error

	return
}

// GetBatchFromDuration 批量查找 耗时分钟
func (obj *_FpTargetMgr) GetBatchFromDuration(durations []int) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`duration` IN (?)", durations).Find(&results).Error

	return
}

// GetFromConcentrateScore 通过concentrate_score获取内容 专注评分
func (obj *_FpTargetMgr) GetFromConcentrateScore(concentrateScore int8) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`concentrate_score` = ?", concentrateScore).Find(&results).Error

	return
}

// GetBatchFromConcentrateScore 批量查找 专注评分
func (obj *_FpTargetMgr) GetBatchFromConcentrateScore(concentrateScores []int8) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`concentrate_score` IN (?)", concentrateScores).Find(&results).Error

	return
}

// GetFromHarvestScore 通过harvest_score获取内容 收获评分
func (obj *_FpTargetMgr) GetFromHarvestScore(harvestScore int8) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`harvest_score` = ?", harvestScore).Find(&results).Error

	return
}

// GetBatchFromHarvestScore 批量查找 收获评分
func (obj *_FpTargetMgr) GetBatchFromHarvestScore(harvestScores []int8) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`harvest_score` IN (?)", harvestScores).Find(&results).Error

	return
}

// GetFromAddTime 通过add_time获取内容 记录添加时间
func (obj *_FpTargetMgr) GetFromAddTime(addTime time.Time) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`add_time` = ?", addTime).Find(&results).Error

	return
}

// GetBatchFromAddTime 批量查找 记录添加时间
func (obj *_FpTargetMgr) GetBatchFromAddTime(addTimes []time.Time) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`add_time` IN (?)", addTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 记录更新时间
func (obj *_FpTargetMgr) GetFromUpdateTime(updateTime time.Time) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 记录更新时间
func (obj *_FpTargetMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FpTargetMgr) FetchByPrimaryKey(id int) (result FpTarget, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpTarget{}).Where("`id` = ?", id).First(&result).Error

	return
}
