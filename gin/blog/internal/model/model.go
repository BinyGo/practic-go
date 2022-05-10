package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/practic-go/gin/blog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Model struct {
	Id         uint32 `gorm:"id" json:"id"`
	CreatedBy  string `gorm:"created_by" json:"created_by"`   // 创建人
	ModifiedBy string `gorm:"modified_by" json:"modified_by"` // 修改人
	CreatedOn  uint32 `gorm:"created_on" json:"created_on"`   // 创建时间
	ModifiedOn uint32 `gorm:"modified_on" json:"modified_on"` // 修改时间
	DeletedOn  uint32 `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      uint8  `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)

	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"

	s = fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)                  // SetMaxOpenConns 设置到数据库的最大打开连接数。
	sqlDB.SetMaxIdleConns(10)                   // SetMaxIdleConns 设置空闲连接池的最大连接数。
	sqlDB.SetConnMaxLifetime(time.Second * 300) // SetConnMaxLifetime 设置连接可以重用的最长时间。

	return db, nil
}
