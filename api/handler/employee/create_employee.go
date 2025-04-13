package employee

import "github.com/gin-gonic/gin"

func CreateEmployeeHandler(ctx *gin.Context) {

	ctx.JSON(201, gin.H{
		"message": "Employee created successfully",
	})

}
