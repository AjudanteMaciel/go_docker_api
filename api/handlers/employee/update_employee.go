package employee

import "github.com/gin-gonic/gin"

func UpdateEmployeeHandler(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Employee updated successfully",
	})

}
