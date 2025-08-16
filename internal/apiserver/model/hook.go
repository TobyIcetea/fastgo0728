package model

import (
	"github.com/TobyIcetea/fastgo/internal/pkg/rid"
	"gorm.io/gorm"
)

// AfterCreate 在创建数据库记录之后生成 postID
func (post *Post) AfterCreate(tx *gorm.DB) error {
	post.PostID = rid.PostID.New(uint64(post.ID))

	return tx.Save(post).Error
}

// AfterCreate 在创建数据库记录之后生成 userID
func (user *User) AfterCreate(tx *gorm.DB) error {
	user.UserID = rid.UserID.New(uint64(user.ID))

	return tx.Save(user).Error
}
