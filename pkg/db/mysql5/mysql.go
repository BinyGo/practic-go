package mysql5

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlDemoCode() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:30306)/practic?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	query := `
		CREATE TABLE users (
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);
	`
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, "58040087", "123456", time.Now())
	if err != nil {
		log.Fatal(err)
	}
	// 获取新插入数据库的用户ID
	fmt.Println("INSERT result:", result)
	userID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("INSERT userID:", userID)

	//查询
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)
	query = `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err = db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, username, password, createdAt)

	//查询多行 保存到结构体
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}
	rows, _ := db.Query(`SELECT id, username, password, created_at FROM users`) //检查错误
	defer rows.Close()
	var users []user
	for rows.Next() {
		var u user
		rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		users = append(users, u)
	}
	fmt.Println(users)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//删除
	result, err = db.Exec(`DELETE FROM users WHERE id = ?`, 8) // 记得检查错误
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DELETE result:", result)
}
