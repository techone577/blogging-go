package impl

import (
	"fmt"

	"github.com/techone577/blogging-go/model"
	"xorm.io/xorm"
)

type tagStorage struct {
	db *xorm.Engine
}

func NewTagStorage(db *xorm.Engine) *tagStorage {
	return &tagStorage{db: db}
}

func (t *tagStorage) QueryEachTagAmount() ([]model.TagAmountInfo, error) {
	var amountInfos []model.TagAmountInfo
	sql := fmt.Sprintf("select tag_id, count(*) as amount from %s where del_flag = 0 group by tag_id order by amount desc", model.TagRelationTableName())
	err := t.db.SQL(sql).Find(&amountInfos)
	return amountInfos, err
}

func (t *tagStorage) QueryByTagIds(ids []int) ([]model.TagInfo, error) {
	var infos []model.TagInfo

	err := t.db.Table(model.TagTableName()).Where("del_flag = ?", 0).In("id", ids).Find(&infos)
	return infos, err
}

// category i dont want to continue fk
func (t *tagStorage) QueryCategories() ([]model.CategoryInfo, error) {
	var cs []model.CategoryInfo

	err := t.db.Table(model.CateTableName()).Where("del_flag = ?", 0).Find(&cs)
	return cs, err
}
