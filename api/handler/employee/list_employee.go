package employee

import "github.com/gin-gonic/gin"

func ListEmployeeHandler(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "List of employees",
	})

}
