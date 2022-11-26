package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FpMemberMgr struct {
	*_BaseMgr
}

// FpMemberMgr open func
func FpMemberMgr(db *gorm.DB) *_FpMemberMgr {
	if db == nil {
		panic(fmt.Errorf("FpMemberMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FpMemberMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fp_member"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FpMemberMgr) GetTableName() string {
	return "fp_member"
}

// Reset 重置gorm会话
func (obj *_FpMemberMgr) Reset() *_FpMemberMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FpMemberMgr) Get() (result FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FpMemberMgr) Gets() (results []*FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FpMemberMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FpMember{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_FpMemberMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 昵称
func (obj *_FpMemberMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithWechat wechat获取 微信
func (obj *_FpMemberMgr) WithWechat(wechat string) Option {
	return optionFunc(func(o *options) { o.query["wechat"] = wechat })
}

// WithAvatar avatar获取 头像
func (obj *_FpMemberMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithIntroduce introduce获取 介绍
func (obj *_FpMemberMgr) WithIntroduce(introduce string) Option {
	return optionFunc(func(o *options) { o.query["introduce"] = introduce })
}

// WithAddTime add_time获取 记录添加时间
func (obj *_FpMemberMgr) WithAddTime(addTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["add_time"] = addTime })
}

// WithUpdateTime update_time获取 记录更新时间
func (obj *_FpMemberMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_FpMemberMgr) GetByOption(opts ...Option) (result FpMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FpMemberMgr) GetByOptions(opts ...Option) (results []*FpMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FpMemberMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FpMember, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where(options.query)
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
func (obj *_FpMemberMgr) GetFromID(id int) (result FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_FpMemberMgr) GetBatchFromID(ids []int) (results []*FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 昵称
func (obj *_FpMemberMgr) GetFromName(name string) (results []*FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 昵称
func (obj *_FpMemberMgr) GetBatchFromName(names []string) (results []*FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromWechat 通过wechat获取内容 微信
func (obj *_FpMemberMgr) GetFromWechat(wechat string) (results []*FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`wechat` = ?", wechat).Find(&results).Error

	return
}

// GetBatchFromWechat 批量查找 微信
func (obj *_FpMemberMgr) GetBatchFromWechat(wechats []string) (results []*FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`wechat` IN (?)", wechats).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容 头像
func (obj *_FpMemberMgr) GetFromAvatar(avatar string) (results []*FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`avatar` = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找 头像
func (obj *_FpMemberMgr) GetBatchFromAvatar(avatars []string) (results []*FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`avatar` IN (?)", avatars).Find(&results).Error

	return
}

// GetFromIntroduce 通过introduce获取内容 介绍
func (obj *_FpMemberMgr) GetFromIntroduce(introduce string) (results []*FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`introduce` = ?", introduce).Find(&results).Error

	return
}

// GetBatchFromIntroduce 批量查找 介绍
func (obj *_FpMemberMgr) GetBatchFromIntroduce(introduces []string) (results []*FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`introduce` IN (?)", introduces).Find(&results).Error

	return
}

// GetFromAddTime 通过add_time获取内容 记录添加时间
func (obj *_FpMemberMgr) GetFromAddTime(addTime time.Time) (results []*FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`add_time` = ?", addTime).Find(&results).Error

	return
}

// GetBatchFromAddTime 批量查找 记录添加时间
func (obj *_FpMemberMgr) GetBatchFromAddTime(addTimes []time.Time) (results []*FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`add_time` IN (?)", addTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 记录更新时间
func (obj *_FpMemberMgr) GetFromUpdateTime(updateTime time.Time) (results []*FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 记录更新时间
func (obj *_FpMemberMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FpMemberMgr) FetchByPrimaryKey(id int) (result FpMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FpMember{}).Where("`id` = ?", id).First(&result).Error

	return
}
