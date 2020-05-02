package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/techone577/blogging-go/global"
	"github.com/techone577/blogging-go/model"
)

func QueryPost(c *gin.Context) {
	var req model.PostQueryRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, fmt.Sprintf("bind request to json failed, err: %v", err))
		return
	}
	if req.PostId == "" {
		c.JSON(400, "post id is empty")
		return
	}
	post, err := global.PostDAO.QueryByPostID(req.PostId)
	if err != nil {
		c.JSON(500, fmt.Sprintf("query post failed, err: %v", err))
		return
	}
	resp := model.PostQueryResponse{
		PostId:       post.PostId,
		Title:        post.Title,
		Category:     post.Category,
		NextPost:     model.PostInfo{Title: "test"},
		PreviousPost: model.PostInfo{Title: "what"},
		HtmlContent:  "hello world",
		AddTime:      "2020-05-02",
		TOC:          "hello",
		StatisticInfo: model.StatisticInfo{
			PageView: 100,
			ReadTime: "2020-05-02",
		},
	}
	c.JSON(200, model.NewUnifiedResponse(true, resp))
}

const (
	queryALL      = "all"
	queryHome     = "home"
	queryCategory = "category"
	queryTag      = "tag"
)

func QueryPostList(c *gin.Context) {
	var req model.PostListQueryRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, fmt.Sprintf("bind request to json failed, err: %v", err))
		return
	}

	switch req.Type {
	case queryALL:
		posts, total, err := global.PostDAO.QueryByPaging(req.Page, req.PageSize, req.ReleaseFlag, req.DelFlag)
		if err != nil {
			c.JSON(400, fmt.Sprintf("query paging failed, err: %v", err))
			return
		}

		var postList []model.HomePagePostInfo
		for _, p := range posts {
			pa, err := global.PostDAO.QueryPassageByPassageId(p.PassageId)
			if err != nil {
				c.JSON(400, fmt.Sprintf("query passage failed, err: %v", err))
				return
			}
			postList = append(postList, model.HomePagePostInfo{
				PostId:        p.PostId,
				Title:         p.Title,
				Summary:       p.Summary,
				AddTime:       "2020-06-02",
				UpdateTime:    "2020-02-04",
				TagList:       []string{"todo"},
				StatisticInfo: model.StatisticInfo{100, ""},
				Category:      p.Category,
				//todo
				FirstImgUrl: pa.PassageId,
			})
		}
		tags, err := QueryTagInfoList()
		if err != nil {
			c.JSON(400, fmt.Sprintf("query tag info list failed, err: %v", err))
			return
		}

		cates, err := QueryCategoryList(true)
		if err != nil {
			c.JSON(400, fmt.Sprintf("query cate info list failed, err: %v", err))
			return
		}

		postListResp := &model.PostListQueryResponse{
			PostList:         postList,
			TagInfoList:      tags,
			CategoryInfoList: cates,
			TotalNum:         total,
			MinId:            0,
		}

		resp := model.NewUnifiedResponse(true, postListResp)
		c.JSON(200, resp)
		return
	case queryHome:
	case queryCategory:
	case queryTag:
	default:
		c.JSON(400, fmt.Sprintf("unknown query type: %s", req.Type))
		return
	}
}

func QueryTagInfoList() ([]model.TagListResp, error) {
	tagAmountInfos, err := global.TagDAO.QueryEachTagAmount()
	if err != nil {
		return nil, err
	}

	ids := []int(nil)
	for _, t := range tagAmountInfos {
		ids = append(ids, t.TagId)
	}
	tagInfos, err := global.TagDAO.QueryByTagIds(ids)
	if err != nil {
		return nil, err
	}
	m := make(map[int]string)
	for _, t := range tagInfos {
		m[t.Id] = t.TagName
	}

	var rs []model.TagListResp

	for _, t := range tagAmountInfos {
		rs = append(rs, model.TagListResp{
			TagName: m[t.TagId],
			TagNum:  t.Amount,
		})
	}
	return rs, nil
}

func QueryCategoryList(isForHome bool) ([]model.CategoryListResp, error) {
	infos, err := global.TagDAO.QueryCategories()
	if err != nil {
		return nil, err
	}
	var rs []model.CategoryListResp
	for _, c := range infos {
		rs = append(rs, model.CategoryListResp{
			Name: c.Name,
			// todo update
			Num:     c.PostNum,
			Summary: c.Summary,
			Url:     c.CoverUrl,
		})
	}
	if len(rs) == 0 {
		return rs, err
	}

	if len(rs) < 5 {
		return rs, err
	}
	if isForHome {
		return rs[0:5], err
	}
	return rs, err
}
