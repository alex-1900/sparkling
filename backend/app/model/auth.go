package model

import "gorm.io/gorm"

// 认证类型
const (
	// AuthTypeEmail 认证类型 email
	AuthTypeEmail = uint8(1 + iota)
	// AuthTypeMobile 认证类型移动电话
	AuthTypeMobile
	// AuthTypeOauth2QQ 认证类型 QQ
	AuthTypeOauth2QQ
	// AuthTypeOauth2Wechat 认证类型微信
	AuthTypeOauth2Wechat
)

// Authorization for User
type Authorization struct {
	gorm.Model
	UserID   uint   `gorm:"not null"`
	Type     uint8  `gorm:"not null;uniqueIndex:idx_type_token"`         // 认证类型
	Token    string `gorm:"size:64;not null;uniqueIndex:idx_type_token"` // 类型唯一标识
	Security string `gorm:"size:64"`                                     // 密钥信息
	Mask     uint   `gorm:"not null"`                                    // 权限 mask
}
