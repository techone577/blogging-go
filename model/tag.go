package model

import "time"

type TagInfo struct {
	Id         int       `xorm:"notnull 'id'" json:"id"`
	TagName    string    `xorm:"notnull 'tag_name'" json:"tagName"`
	DelFlag    string    `xorm:"notnull 'del_flag'" json:"delFlag"`
	AddTime    time.Time `xorm:"notnull 'add_time'" json:"addTime"`
	UpdateTime time.Time `xorm:"notnull 'update_time'" json:"updateTime"`
}

func TagTableName() string {
	return "tag"
}

type TagListResp struct {
	TagName string `json:"tagName"`
	TagNum  int    `json:"tagNum"`
}

type TagAmountInfo struct {
	TagId  int `xorm:"notnull 'tag_id'" json:"tagId"`
	Amount int `xorm:"notnull 'amount'" json:"amount"`
}

type TagRelation struct {
}

func TagRelationTableName() string {
	return "tag_relation"
}
