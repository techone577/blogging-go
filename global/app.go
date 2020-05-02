package global

import (
	"github.com/techone577/blogging-go/dao"
	"xorm.io/xorm"
)

var (
	DB *xorm.Engine

	PostDAO dao.PostStorage
	TagDAO  dao.TagStorage
)
