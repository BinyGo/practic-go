package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/practic-go/gin/blog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Model struct {
	ID         uint32 `gorm:"id" json:"id"`
	CreatedBy  string `gorm:"created_by" json:"created_by"`   // 创建人
	ModifiedBy string `gorm:"modified_by" json:"modified_by"` // 修改人
	CreatedOn  uint32 `gorm:"created_on" json:"created_on"`   // 创建时间
	ModifiedOn uint32 `gorm:"modified_on" json:"modified_on"` // 修改时间
	DeletedOn  uint32 `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      uint8  `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	// 1. 初始化Jaeger
	// closer, err := tracer_gorm.InitJaeger()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer closer.Close()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Warn, // Log level
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

	//db, err := gorm.Open(mysql.Open(s), &gorm.Config{Logger: newLogger})
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}
	// 3. 最重要的一步，使用我们定义的插件
	// _ = db.Use(&tracer_gorm.OpentracingPlugin{})
	// // 4. 生成新的Span - 注意将span结束掉，不然无法发送对应的结果
	// span := opentracing.StartSpan("gormTracing unit test")
	// defer span.Finish()
	// // 5. 把生成的Root Span写入到Context上下文，获取一个子Context
	// // 通常在Web项目中，Root Span由中间件生成
	// ctx := opentracing.ContextWithSpan(context.Background(), span)
	// db.WithContext(ctx)

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
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Create().Replace("gorm:before_create", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Update().Replace("gorm:before_update", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	return db, nil
}

func updateTimeStampForCreateCallback(db *gorm.DB) {

	if db.Error == nil {
		if createTimeField, ok := db.Statement.Schema.FieldsByName["CreatedOn"]; ok {
			if !createTimeField.NotNull {
				_ = createTimeField.AutoCreateTime
			}
		}

		if createTimeField, ok := db.Statement.Schema.FieldsByName["ModifiedOn"]; ok {
			if !createTimeField.NotNull {
				_ = createTimeField.AutoUpdateTime
			}
		}
		//db.Statement.SetColumn("CreatedOn", time.Now().Unix())

	}
}

func updateTimeStampForUpdateCallback(db *gorm.DB) {
	if _, ok := db.Get("gorm:update_column"); !ok {
		_ = db.Set("ModifiedOn", time.Now().Unix())
	}
	//db.Statement.SetColumn("ModifiedOn", time.Now().Unix())
}

func deleteCallback(db *gorm.DB) {
	if db.Error == nil {
		if db.Statement.Schema != nil {
			db.Statement.SQL.Grow(100)
			deleteField := db.Statement.Schema.LookUpField("DeletedOn")
			if !db.Statement.Unscoped && deleteField != nil {
				//Soft Delete
				if db.Statement.SQL.String() == "" {
					nowTime := time.Now().Unix()
					db.Statement.AddClause(
						clause.Set{{
							Column: clause.Column{Name: deleteField.DBName},
							Value:  nowTime,
						}},
					)
					db.Statement.AddClauseIfNotExists(clause.Update{})
					db.Statement.Build("UPDATE", "SET", "WHERE")
				}
			} else {
				//Delete
				if db.Statement.SQL.String() == "" {
					db.Statement.AddClauseIfNotExists(clause.Delete{})
					db.Statement.AddClauseIfNotExists(clause.From{})
					db.Statement.Build("DELETE", "FROM", "WHERE")
				}
			}
			fmt.Println(db.Statement.SQL.String())
			fmt.Println(db.Statement.Vars)
			//Must Need WHERE
			if _, ok := db.Statement.Clauses["WHERE"]; !db.AllowGlobalUpdate && !ok {
				db.AddError(gorm.ErrMissingWhereClause)
				return
			}
			db.Exec(db.Statement.SQL.String(), db.Statement.Vars...)
		}
	}
}
