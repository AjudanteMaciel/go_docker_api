package router

import (
	"go_docker_api/api/handler/employee"
	"go_docker_api/api/handler/handlerutil"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/employees", employee.ListEmployeeHandler)
		v1.GET("/employee/:id", employee.ShowEmployeeHandler)
		v1.POST("/employee", employee.CreateEmployeeHandler)
		v1.PUT("/employee/:id", employee.UpdateEmployeeHandler)
		v1.DELETE("/employee/:id", employee.DeleteEmployeeHandler)
	}
	{
		v1.GET("/ping", handlerutil.PingHandler)
	}
	// I KNOW THIS IS DUMB BUT I AM JUST TESTING THE DROP DATABASE FUNCTIONALITY
	{
		v1.GET("/drop_database", handlerutil.DropDatabaseHandler)
	}

}
