package model

import (
	"gorm.io/gorm"
)

const (
	UserStatusNormal int = iota + 1
	UserStatusClosed
)

type User struct {
	*Model
	Nickname    string `json:"nickname,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	Salt        string `json:"salt,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Status      int    `json:"status,omitempty"`
	IsStaff     bool   `json:"is_staff,omitempty"`
	IsSuperuser bool   `json:"is_superuser,omitempty"`
}

func (u *User) Create(db *gorm.DB) (*User, error) {
	err := db.Create(&u).Error
	return u, err
}

func (u *User) Update(db *gorm.DB) error {
	return db.Model(&User{}).
		Where("id = ? AND is_deleted = ?", u.Model.ID, 0).
		Save(u).
		Error
}

func (u *User) Get(db *gorm.DB) (*User, error) {
	var user User
	if u.Model != nil && u.Model.ID > 0 {
		db = db.Where("id = ? AND is_deleted = ?", u.Model.ID, 0)
	} else {
		db = db.Where("username = ? AND is_deleted = ?", u.Username, 0)
	}
	err := db.First(&user).Error
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (u *User) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*User, error) {
	var users []*User
	var err error
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}
	for k, v := range *conditions {
		if k == "ORDER" {
			db = db.Order(v)
		} else {
			db = db.Where(k, v)
		}
	}
	if err = db.Where("is_deleted = ?", 0).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
