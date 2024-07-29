package handler

import (
	"github.com/luizhirasawa/Go-api.git/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("Go-API: ")
	db = config.GetPsql()

	//create operations
}
