package dao

import (
	"github.com/techone577/blogging-go/global"
	"github.com/techone577/blogging-go/model"
)

type PostStorage interface {
	QueryByPostID(id string) (*model.PostInfo, error)
}

type postStorage struct{}

func NewPostStorage() PostStorage {
	return &postStorage{}
}

func (p *postStorage) QueryByPostID(id string) (*model.PostInfo, error) {
	var post model.PostInfo
	_, err := global.DB.Table(model.PostTableName()).Where("post_id = ? and del_flag = ? and release_flag = ?", id, 0, 1).Get(&post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
