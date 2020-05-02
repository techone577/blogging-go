package http

import (
	"github.com/gin-gonic/gin"
	"github.com/techone577/blogging-go/http/handler"
)

func Route(e *gin.Engine) error {
	//e.Use(cors.New(cors.Config{
	//	AllowMethods:     []string{"GET", "POST"},
	//	AllowHeaders:     []string{"Content-Type"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))

	e.GET("/home", handler.HomePage)
	e.GET("/about", handler.AboutPage)
	e.GET("/postlist", handler.PostList)
	e.GET("/tagShow", handler.TagShow)
	e.GET("/categoryShow", handler.CategoryShow)
	e.GET("/search", handler.Search)

	e.POST("/queryPost", handler.QueryPost)
	e.POST("/queryPostList", handler.QueryPostList)
	return nil
}
