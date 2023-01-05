package core

import "algo-archive/internal/model"

// UserManageService User management service
type UserManageService interface {
	GetUserByID(id int64) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUsersByIDs(id [][]int64) ([]*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) error
}
