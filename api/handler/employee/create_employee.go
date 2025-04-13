package employee

import (
	"go_docker_api/api/handler"
	"go_docker_api/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEmployeeHandler(ctx *gin.Context) {
	request := CreateEmployeeRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		handler.Logger.Errorf("validation error: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	employee := schemas.Employee{
		Name:    request.Name,
		Email:   request.Email,
		RoleID:  request.RoleID,
		StateID: request.StateID,
	}

	if err := handler.Db.Create(&employee).Error; err != nil {
		handler.Logger.Errorf("error creating opening: %v", err.Error())
		handler.SendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	handler.SendSuccess(ctx, "create employee", employee)
}
