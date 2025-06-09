package handler

import (
	"encoding/json"
	"net/http"
	"register/internal/repository"
	"register/internal/service"
)

type RegisterHandler struct {
	userRepo repository.UserRepository
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// http.Handler는 ServeHTTP를 구현해야함
func NewRegisterHandler(userRepo repository.UserRepository) http.Handler {
	return &RegisterHandler{userRepo: userRepo}
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	encryptPassword, err := service.EncryptPassword(req.Password)
	if err != nil {
		http.Error(w, "Password Encrypt Error", http.StatusInternalServerError)
		return
	}

	if err := h.userRepo.CreateUser(req.Username, encryptPassword); err != nil {
		http.Error(w, "DB 에러", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
