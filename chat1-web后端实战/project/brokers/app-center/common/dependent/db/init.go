package db

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // 导入postgres
	_ "github.com/lib/pq"
	"time"
)

var DB *gorm.DB
var PQDB *sql.DB

func init() {
	// 初始化数据库orm连接

	dataSource := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
		"localhost",
		"postgres",
		"test",
		"disable",
		"123",
	)
	db, err := gorm.Open("postgres",
		dataSource,
	)
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetConnMaxLifetime(2 * time.Second)
	db.DB().SetMaxIdleConns(30)
	if err != nil {
		panic(err)
	} else {
		DB = db
	}
	if e := DB.DB().Ping(); e != nil {
		panic(e)
	}
	// 初始化pq驱动,用于CopyIn
	PQDB, err = sql.Open("postgres", dataSource)
	PQDB.SetMaxIdleConns(1)
	if err != nil {
		panic(err)
	}
}
