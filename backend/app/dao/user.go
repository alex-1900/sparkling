package dao

import (
	"github.com/alex-1900/sparkling/app/model"
	"github.com/alex-1900/sparkling/app/service"
)

type User struct {
	userModel *model.User
}

// Save 保存用户数据
func (u *User) Save() (*User, error) {
	main := service.GetDB().Main()
	if err := main.Save(u.userModel).Error; err != nil {
		return nil, err
	}
	return &User{userModel: u.userModel}, nil
}
