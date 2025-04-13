package handlerutil

import (
	"go_docker_api/configs"

	"github.com/gin-gonic/gin"
)

// I KNOW THIS IS DUMB BUT I AM JUST TESTING THE DROP DATABASE FUNCTIONALITY
func DropDatabaseHandler(ctx *gin.Context) {

	configs.GetPostgres().Migrator().DropTable(configs.SchemasList...)

}
