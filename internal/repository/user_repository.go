package repository

import "database/sql"

type UserRepository interface {
	CreateUser(username, encryptPassword string) error
}

type userRepository struct {
	db *sql.DB
}

const (
	insertUserQuery = `INSERT INTO users (username, password) VALUES (?, ?)`
)

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(username, encryptPassword string) error {
	_, err := r.db.Exec(insertUserQuery, username, encryptPassword)
	return err
}

// 인터페이스의 모든 함수를 구현한 구현체(만족)는
// 명시적으로 인터페이스를 상속하지 않았어도
// 인터페이스 타입처럼 사용 가능함
//
// (r *userRepository) CreateUser 이 됨으로써
// userRepository가 UserRepository를 만족하고 해당 Type처럼 사용 가능
