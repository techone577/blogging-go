package handler

import (
	"github.com/gin-gonic/gin"
)

func AboutPage(c *gin.Context) {
	c.HTML(200, "about.html", "")
}

func HomePage(c *gin.Context) {
	c.HTML(200, "homePage.html", "")
}

func PostList(c *gin.Context) {
	c.HTML(200, "postlist.html", "")
}

func TagShow(c *gin.Context) {
	c.HTML(200, "tagShow.html", "")
}

func CategoryShow(c *gin.Context) {
	c.HTML(200, "categoryShow.html", "")
}

func Search(c *gin.Context) {
	c.HTML(200, "search.html", "")
}
