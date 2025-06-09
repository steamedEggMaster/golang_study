package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	env := os.Getenv("ENV")
	if env == "" || env == "local" {
		err := godotenv.Load() // .env파일 로드하기
		if err != nil {
			log.Fatal("env 파일 로드 불가 : ", err)
		}
	}
}
