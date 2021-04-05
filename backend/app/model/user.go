package model

import "gorm.io/gorm"

// 用户状态
const (
	// UserStatusActive 用户状态启用
	UserStatusActive = uint8(1 + iota)
	// UserStatusFrozen 用户状态冻结
	UserStatusFrozen
	// UserStatusInactive 用户状态禁用
	UserStatusInactive
)

// 用户性别
const (
	// UserGenderMale 用户性别 男
	UserGenderMale = uint8(1 + iota)
	// UserGenderFemale 用户性别 女
	UserGenderFemale
	// UserGenderSecret 用户性别 保密
	UserGenderSecret
)

// User model
type User struct {
	gorm.Model
	Nickname      string        `gorm:"size:64;not null;index:idx_nickname"` // 昵称
	Gender        uint8         `gorm:"not null"`                            // 性别
	Bio           string        `gorm:"type:text"`                           // 简介
	Location      string        `gorm:"size:128"`                            // 地区
	Email         string        `gorm:"size:64"`                             // 公开的 Email
	Status        uint8         `gorm:"not null"`                            // 状态
	Authorization Authorization `gorm:"foreignkey:UserID"`
}
