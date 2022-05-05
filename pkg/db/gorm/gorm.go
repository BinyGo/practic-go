package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var sqlDB *sql.DB

func initDB() (*gorm.DB, *sql.DB) {

	dsn := "root:root@tcp(127.0.0.1:30306)/practic?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)                  // SetMaxOpenConns 设置到数据库的最大打开连接数。
	sqlDB.SetMaxIdleConns(10)                   // SetMaxIdleConns 设置空闲连接池的最大连接数。
	sqlDB.SetConnMaxLifetime(time.Second * 300) // SetConnMaxLifetime 设置连接可以重用的最长时间。

	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}
	return db, sqlDB

}

type User struct {
	Id        int64     `gorm:"column:id;primary_key"`
	UserName  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func main() {
	db, sqlDB = initDB()
	user := &User{}
	db.Select("id, username, password, created_at").Find(user)
	fmt.Println(user)

	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err := sqlDB.QueryRow(query, 3).Scan(&user.Id, &user.UserName, &user.Password, &user.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}
