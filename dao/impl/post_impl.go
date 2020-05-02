package impl

import (
	"fmt"
	"github.com/techone577/blogging-go/model"
	"xorm.io/xorm"
)

type postStorage struct {
	db *xorm.Engine
}

func NewPostStorage(db *xorm.Engine) *postStorage {
	return &postStorage{db: db}
}

func (p *postStorage) QueryByPostID(id string) (*model.PostInfo, error) {
	var post model.PostInfo
	_, err := p.db.Table(model.PostTableName()).Where("post_id = ? and del_flag = ? and release_flag = ?", id, 0, 1).Get(&post)
	return &post, err
}

func (p *postStorage) QueryPreviousPost(pkId int) (*model.PostInfo, error) {
	var post model.PostInfo

	innerSql := fmt.Sprintf("select min(id) from %s where id > %d", model.PostTableName(), pkId)
	sql := fmt.Sprintf("select * from %s where del_flag = 0 and release_flag = 1 and id in (%s)", model.PostTableName(), innerSql)

	_, err := p.db.SQL(sql).Get(&post)
	return &post, err
}

func (p *postStorage) QueryNextPost(pkId int) (*model.PostInfo, error) {
	var post model.PostInfo

	innerSql := fmt.Sprintf("select max(id) from %s where id < %d", model.PostTableName(), pkId)
	sql := fmt.Sprintf("select * from %s where del_flag = 0 and release_flag = 1 and id in (%s)", model.PostTableName(), innerSql)

	_, err := p.db.SQL(sql).Get(&post)
	return &post, err
}

func (p *postStorage) QueryByPaging(page, pageSize, releaseFlag, delFlag int) ([]model.PostInfo, int64, error) {
	var posts []model.PostInfo

	s := p.db.Table(model.PostTableName()).Where("release_flag = ? and del_flag = ?", releaseFlag, delFlag)
	total, err := s.Count()
	if err != nil {
		return nil, 0, err
	}
	err = s.Limit(pageSize, (page-1)*pageSize).Desc("id").Find(&posts)
	return posts, total, err
}

func (p *postStorage) QueryPassageByPassageId(id string) (model.PassageInfo, error) {
	var pa model.PassageInfo

	_, err := p.db.Table(model.PassageTableName()).Where("del_flag = ? and passage_id = ?", 0, id).Get(&pa)
	return pa, err
}
