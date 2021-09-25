// Package app 应用封装层
package app

import (
	"gitee.com/itse/personal-work/app/config"
	"gitee.com/itse/personal-work/app/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// InitGorm 初始化数据库引擎
func InitGorm() (*gorm.DB, error) {
	c := config.C.GORM
	db, err := NewDB()
	if err != nil {
		return nil, err
	}

	if c.EnableAutoMigrate {
		err = AutoMigrate(db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

// NewDB 创建 DB 实例
func NewDB() (*gorm.DB, error) {
	c := config.C.DB
	cGorm := config.C.GORM

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             2 * time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Silent,   // 日志级别
			IgnoreRecordNotFoundError: true,            // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,           // 禁用彩色打印
		},
	)

	db, err := gorm.Open(sqlite.Open(c.DSN()), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	if cGorm.Debug {
		db = db.Debug()
	}

	sql, _ := db.DB()
	err = sql.Ping()
	if err != nil {
		return nil, err
	}

	sql.SetMaxIdleConns(cGorm.MaxIdleConns)
	sql.SetMaxOpenConns(cGorm.MaxOpenConns)
	sql.SetConnMaxLifetime(time.Duration(cGorm.MaxLifetime) * time.Second)

	return db, nil
}

// AutoMigrate 自动映射数据表
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{}, &entity.Role{}, &entity.Title{})
}
