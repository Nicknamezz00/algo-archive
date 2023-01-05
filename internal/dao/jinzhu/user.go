package jinzhu

import (
	"algo-archive/internal/core"
	"algo-archive/internal/model"
	"gorm.io/gorm"
)

type userManageServant struct {
	db *gorm.DB
}

var (
	_ core.UserManageService = (*userManageServant)(nil)
)

func newUserManageService(db *gorm.DB) core.UserManageService {
	return &userManageServant{db: db}
}

func (s *userManageServant) GetUserByID(id int64) (*model.User, error) {
	user := &model.User{
		Model: &model.Model{ID: id},
	}
	return user.Get(s.db)
}

func (s *userManageServant) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{
		Username: username,
	}
	return user.Get(s.db)
}

func (s *userManageServant) GetUsersByIDs(ids [][]int64) ([]*model.User, error) {
	user := &model.User{}
	return user.List(s.db, &model.ConditionsT{
		"id IN ?": ids,
	}, 0, 0)
}

func (s *userManageServant) CreateUser(user *model.User) (*model.User, error) {
	return user.Create(s.db)
}

func (s *userManageServant) UpdateUser(user *model.User) error {
	return user.Update(s.db)
}
