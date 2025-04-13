package employee

import "github.com/gin-gonic/gin"

func DeleteEmployeeHandler(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Employee deleted successfully",
	})

}
