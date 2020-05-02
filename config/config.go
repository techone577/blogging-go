package config

import (
	"fmt"
	"time"

	"github.com/techone577/blogging-go/dao/impl"
	"github.com/techone577/blogging-go/global"
	"xorm.io/xorm"

	_ "github.com/go-sql-driver/mysql"
)

func MustConfigure() {
	mustConfigureMysql()
	mustConfigureStorage()
}

func mustConfigureMysql() {
	global.DB = mustSetupMySQL()

}

func mustConfigureStorage() {
	global.PostDAO = impl.NewPostStorage(global.DB)
	global.TagDAO = impl.NewTagStorage(global.DB)
}

func mustSetupMySQL() *xorm.Engine {
	engine, err := newEngine()
	if err != nil {
		panic(fmt.Errorf("fail to init DB: %v", err))
	}
	return engine
}

func newEngine() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", "root:techone577@tcp(39.108.184.185:4306)/blogging")
	if err != nil {
		return nil, fmt.Errorf("db: fail to create xorm engine: %v", err)
	}

	err = engine.Ping()
	if err != nil {
		engine.Close()
		return nil, fmt.Errorf("db: fail to ping database: %v", err)
	}

	// Store UTC in DB tables instead of server local time
	defaultZone, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	engine.SetTZDatabase(defaultZone)
	engine.SetMaxOpenConns(30)
	engine.SetMaxIdleConns(30)
	engine.SetConnMaxLifetime(time.Duration(3600) * time.Second)
	return engine, err
}
