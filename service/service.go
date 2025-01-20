package service

import "database/sql"

type Service struct {
	User UserService
}

func NewService(db *sql.DB) *Service {
	return &Service{
		User: newUserService(db),
	}
}
