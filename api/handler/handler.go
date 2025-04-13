package handler

import (
	"go_docker_api/configs"

	"gorm.io/gorm"
)

var (
	Logger *configs.Logger
	Db     *gorm.DB
)

func InitializeHandler() {
	Logger = configs.GetLogger("handler")
	Db = configs.GetPostgres()
}
