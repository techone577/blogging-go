package dao

import (
	"github.com/techone577/blogging-go/model"
)

type PostStorage interface {
	QueryByPostID(id string) (*model.PostInfo, error)
	QueryPreviousPost(pkId int) (*model.PostInfo, error)
	QueryNextPost(pkId int) (*model.PostInfo, error)
	QueryByPaging(page, pageSize, releaseFlag, delFlag int) ([]model.PostInfo, int64, error)

	QueryPassageByPassageId(id string) (model.PassageInfo, error)
}

type TagStorage interface {
	QueryEachTagAmount() ([]model.TagAmountInfo, error)
	QueryByTagIds(ids []int) ([]model.TagInfo, error)
	QueryCategories() ([]model.CategoryInfo, error)
}
