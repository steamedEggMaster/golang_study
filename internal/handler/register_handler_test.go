package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeUserRepo struct {
	createCalled bool
	lastUsername string
	lastPassword string
	returnError  error
}

func (f *fakeUserRepo) CreateUser(username, password string) error {
	f.createCalled = true
	f.lastUsername = username
	f.lastPassword = password
	return f.returnError
}

func TestRegisterHandler_ServeHTTP_Success(t *testing.T) {
	repo := &fakeUserRepo{}
	h := NewRegisterHandler(repo)

	body := RegisterRequest{
		Username: "testuser",
		Password: "password123",
	}
	bodyBytes, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("예상 상태 코드 %d, 실제 %d", http.StatusCreated, rec.Code)
	}
	if !repo.createCalled {
		t.Error("CreateUser가 호출되지 않음")
	}
}
