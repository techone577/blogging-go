package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/techone577/blogging-go/config"

	bhttp "github.com/techone577/blogging-go/http"
)

func main() {
	flag.Parse()

	config.MustConfigure()

	r := gin.Default()
	r.LoadHTMLGlob("http/html/*")
	r.Static("/static", "./http/static")
	r.Static("/fonts", "./http/static/fonts")
	r.Static("/css", "./http/static/css")
	r.Static("/js", "./http/static/js")
	r.Static("/images", "./http/static/images")
	r.Static("/layui", "./http/static/layui")

	if err := bhttp.Route(r); err != nil {
		log.Fatalf("route failed, err: %v", err)
	}
	if err := r.Run(); err != nil {
		log.Fatalf("hello world")
	}
}
