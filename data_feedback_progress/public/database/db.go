package database

import (
	"context"
	"data_feedback_progress/public/config"
	zap2 "data_feedback_progress/public/zap"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

var (
	fpMysqlDb   *gorm.DB
	fpMysqlOnce sync.Once
)

type MysqlLogger struct {
	*zap.SugaredLogger
}

func (t *MysqlLogger) LogMode(logger.LogLevel) logger.Interface { return nil }
func (t *MysqlLogger) Info(c context.Context, s string, i ...interface{}) {
}
func (t *MysqlLogger) Warn(c context.Context, s string, i ...interface{}) {
}
func (t *MysqlLogger) Error(c context.Context, s string, i ...interface{}) {
}
func (t *MysqlLogger) Trace(c context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	s, r := fc()
	t.Debugf("%dms, sql: %s rowsAffected: %d", time.Now().Sub(begin).Milliseconds(), s, r)
}

func GormFp() *gorm.DB {
	if fpMysqlDb == nil {
		fpMysqlOnce.Do(func() {
			fpMysqlInit()
		})
	}
	return fpMysqlDb
}

func fpMysqlInit() {
	logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(config.FpMysqlDsn), &gorm.Config{
		Logger: &MysqlLogger{
			SugaredLogger: zap2.Zap(),
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDb, _ := db.DB()
	sqlDb.SetConnMaxLifetime(time.Second * 3600 * 6)
	sqlDb.SetMaxOpenConns(config.FpMysqlMaxConn)
	sqlDb.SetMaxIdleConns(10)
	zap2.Zap().Infof("szb mysql max conn: %d", config.FpMysqlMaxConn)

	fpMysqlDb = db
	return
}
