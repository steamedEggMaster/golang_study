package main

import (
	"log"
	"net/http"
	"register/database"
	"register/internal/handler"
	"register/internal/repository"
)

func main() {
	db := database.NewMySQL()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	http.Handle("/register", handler.NewRegisterHandler(userRepo))

	log.Println("서버 시작 : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
