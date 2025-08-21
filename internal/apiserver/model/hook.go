package model

import (
	"github.com/TobyIcetea/fastgo/internal/pkg/rid"
	"github.com/TobyIcetea/fastgo/pkg/auth"

	"gorm.io/gorm"
)

// AfterCreate 在创建数据库记录之后生成 postID
func (post *Post) AfterCreate(tx *gorm.DB) error {
	post.PostID = rid.PostID.New(uint64(post.ID))

	return tx.Save(post).Error
}

// BeforeCreate 在创建数据库记录之前加密明文密码
func (m *User) BeforeCreate(tx *gorm.DB) error {
	// Encrypt the user password.
	var err error
	m.Password, err = auth.Encrypt(m.Password)
	if err != nil {
		return err
	}
	return nil
}

// AfterCreate 在创建数据库记录之后生成 userID
func (user *User) AfterCreate(tx *gorm.DB) error {
	user.UserID = rid.UserID.New(uint64(user.ID))

	return tx.Save(user).Error
}
