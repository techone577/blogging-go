package model

import "time"

type CategoryListResp struct {
	Name    string `json:"name"`
	Num     int    `json:"num"`
	Summary string `json:"summary"`
	Url     string `json:"url"`
}

type CategoryInfo struct {
	Id         int       `xorm:"notnull 'id'" json:"id"`
	Name       string    `xorm:"notnull 'name'" json:"name"`
	CoverUrl   string    `xorm:"notnull 'cover_url'" json:"coverUrl"`
	Summary    string    `xorm:"notnull 'summary'" json:"summary"`
	PostNum    int       `xorm:"notnull 'post_num'" json:"postNum"`
	AddTime    time.Time `xorm:"notnull 'add_time'" json:"add_time"`
	UpdateTime time.Time `xorm:"notnull 'update_time'" json:"update_time"`
	DelFlag    int       `xorm:"notnull 'del_flag'" json:"del_flag"`
}

func CateTableName() string {
	return "category"
}
