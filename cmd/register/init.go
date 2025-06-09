package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	env := os.Getenv("ENV")
	envFile := fmt.Sprintf("config/env/%s.env", env)

	err := godotenv.Load(envFile) // .env파일 로드하기
	if err != nil {
		log.Fatalf("%s.env 파일 로드 불가 : %v", env, err)
	}
}
