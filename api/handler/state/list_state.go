package state

import (
	"go_docker_api/api/handler"
	"go_docker_api/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListStateHandler(ctx *gin.Context) {

	states := []schemas.State{}

	if err := handler.Db.Find(&states).Error; err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	handler.SendSuccess(ctx, "list-states", states)

}
