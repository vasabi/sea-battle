package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var VarCreateBattleField = CreateBattleField

// GameRoutesGroup - game routes.
func GameRoutesGroup(route *gin.RouterGroup) {
	route.POST("/battle/create-matrix", VarCreateBattleField)
}

//  TODO: models
// CreateBattleField creates a sea battle game field.
// @Description Создать поле для игры в морской бой.
// @Tags battle
// @Accept  json
// @Produce  json
// @Success 200 {object} swagger.WithError
// @Failure 500 {object} swagger.WithError
// @Router /v1/rest/battle/create-matrix [post]
func CreateBattleField(c *gin.Context) {
	//TODO: handler
	c.JSON(http.StatusOK, WrapResponse(true, ""))
}
