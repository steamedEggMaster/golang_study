package database

import (
	"database/sql"
	"log"
	"register/config"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL() *sql.DB {
	cfg := config.LoadDBConfig()

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		log.Fatalf("DB 연결 실패: %v", err)
	}
	log.Print("DB 연결 성공")

	// 연결 확인
	if err := db.Ping(); err != nil {
		log.Fatalf("DB Ping 실패: %v", err)
	}
	log.Print("DB Ping 성공")

	return db
}
