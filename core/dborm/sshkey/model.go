package sshkey

import (
	"tdp-cloud/core/dborm"
)

// 添加密钥

type CreateParam struct {
	UserId      uint   `json:"userId"`
	PublicKey   string `json:"publicKey" binding:"required"`
	PrivateKey  string `json:"privateKey" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func Create(post *CreateParam) error {

	result := dborm.Db.Create(&dborm.Sshkey{
		UserId:      post.UserId,
		PublicKey:   post.PublicKey,
		PrivateKey:  post.PrivateKey,
		Description: post.Description,
	})

	return result.Error

}

// 更新密钥

type UpdateParam struct {
	Id          uint   `json:"id"  binding:"required"`
	UserId      uint   `json:"userId" binding:"required"`
	PublicKey   string `json:"publicKey" binding:"required"`
	PrivateKey  string `json:"privateKey" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func Update(post *UpdateParam) error {

	result := dborm.Db.Model(&dborm.Sshkey{}).
		Where("id = ? AND user_id = ?", post.Id, post.UserId).
		Updates(dborm.Sshkey{
			PublicKey:   post.PublicKey,
			PrivateKey:  post.PrivateKey,
			Description: post.Description,
		})

	return result.Error

}

// 获取密钥列表

func FetchAll(userId uint) ([]*dborm.Sshkey, error) {

	var items []*dborm.Sshkey

	result := dborm.Db.Find(&items, "user_id = ?", userId)

	return items, result.Error

}

// 获取密钥

func Fetch(id, userId uint) (*dborm.Sshkey, error) {

	var item *dborm.Sshkey

	result := dborm.Db.First(&item, "id = ? AND user_id = ?", id, userId)

	return item, result.Error

}

// 删除密钥

func Delete(id, userId uint) error {

	var item *dborm.Sshkey

	result := dborm.Db.Delete(&item, "id = ? AND user_id = ?", id, userId)

	return result.Error

}