package employee

import "github.com/gin-gonic/gin"

func ShowEmployeeHandler(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Show employee details",
	})

}
