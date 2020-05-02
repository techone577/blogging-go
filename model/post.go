package model

import "time"

type PostQueryRequest struct {
	PostId string `json:"postId"`
}

type PostQueryResponse struct {
	PostId        string        `json:"postId"`
	Title         string        `json:"title"`
	Category      string        `json:"category"`
	NextPost      PostInfo      `json:"nextPost"`
	PreviousPost  PostInfo      `json:"previousPost"`
	HtmlContent   string        `json:"htmlContent"`
	AddTime       string        `json:"addTime"`
	TOC           string        `json:"toc"`
	StatisticInfo StatisticInfo `json:"statisticInfo"`
}

type PostListQueryRequest struct {
	Page        int    `json:"page"`
	PageSize    int    `json:"pageSize"`
	Type        string `json:"type"`
	TypeValue   string `json:"typeValue"`
	ReleaseFlag int    `json:"releaseFlag"`
	DelFlag     int    `json:"delFlag"`
}

type PostListQueryResponse struct {
	PostList         []HomePagePostInfo `json:"postList"`
	TagInfoList      []TagListResp      `json:"tagInfoList"`
	CategoryInfoList []CategoryListResp `json:"categoryInfoList"`
	TotalNum         int64              `json:"totalNum"`
	MinId            int                `json:"minId"`
}

type HomePagePostInfo struct {
	PostId        string        `json:"postId"`
	Title         string        `json:"title"`
	Summary       string        `json:"summary"`
	AddTime       string        `json:"addTime"`
	UpdateTime    string        `json:"updateTime"`
	TagList       []string      `json:"tagList"`
	StatisticInfo StatisticInfo `json:"statisticInfo"`
	Category      string        `json:"category"`
	FirstImgUrl   string        `json:"firstImgUrl"`
}

type PostInfo struct {
	ID          int       `xorm:"autoincr notnull pk 'id'" json:"id"`
	PostId      string    `xorm:"varchar(32) notnull 'post_id'" json:"postId"`
	Title       string    `xorm:"varchar(128) notnull 'title'" json:"title"`
	Category    string    `xorm:"varchar(32) notnull 'category'" json:"category"`
	PassageId   string    `xorm:"varchar(32) notnull 'passage_id'" json:"passageId"`
	ReleaseFlag int       `xorm:"notnull 'release_flag'" json:"releaseFlag"`
	DelFlag     int       `xorm:"notnull 'del_flag'" json:"delFlag"`
	AddTime     time.Time `xorm:"notnull 'add_time'" json:"addTime"`
	UpdateTime  time.Time `xorm:"notnull 'update_time'" json:"updateTime"`
	Summary     string    `xorm:"varchar(255) notnull 'summary'" json:"summary"`
}

func PostTableName() string {
	return "post_info"
}

type PassageInfo struct {
	ID         int       `xorm:"autoincr notnull pk 'id'" json:"id"`
	PassageId  string    `xorm:"varchar(32) notnull 'passage_id'" json:"passageId"`
	Content    string    `xorm:"notnull 'content'" json:"content"`
	DelFlag    int       `xorm:"notnull 'del_flag'" json:"delFlag"`
	AddTime    time.Time `xorm:"notnull 'add_time'" json:"addTime"`
	UpdateTime time.Time `xorm:"notnull 'update_time'" json:"updateTime"`
}

func PassageTableName() string {
	return "passage"
}

type StatisticInfo struct {
	PageView int
	ReadTime string
}
