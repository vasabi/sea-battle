package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	VarCreateBattleField = CreateBattleField
	VarCreateFleet       = CreateFleet
	VarShot              = Shot
	VarClear             = ClearBattleField
	VarState             = GetState
)

// GameRoutesGroup - game routes.
func GameRoutesGroup(route *gin.RouterGroup) {
	route.POST("/battle/create-matrix", VarCreateBattleField)
	route.POST("/battle/ship", VarCreateFleet)
	route.POST("/battle/shot", VarShot)
	route.POST("/battle/clear", VarClear)
	route.GET("/battle/state", VarState)
}

// CreateBattleField creates a sea battle game field.
// @Description Создать поле для игры в морской бой.
// @Tags battle
// @Accept  json
// @Produce  json
// @Param parameters body GameFieldReq true "JSON Body with parameters"
// @Success 200 {string} string "message"
// @Failure 400 {string} string	"error"
// @Router /rest/v1/battle/create-matrix [post]
func CreateBattleField(c *gin.Context) {
	var reqBody GameFieldReq

	if err := c.BindJSON(&reqBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, reqBody.CreateGameField())
}

// CreateFleet creates a player's fleet.
// @Description Разместить корабли на игровом поле.
// @Tags battle
// @Accept  json
// @Produce  json
// @Param parameters body CreateFleetReq true "JSON Body with parameters"
// @Success 200 {string} string "message"
// @Failure 400 {string} string	"error"
// @Router /rest/v1/battle/ship [post]
func CreateFleet(c *gin.Context) {
	var reqBody CreateFleetReq

	if err := c.BindJSON(&reqBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := reqBody.CreateFleet(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "fleet crated")
}

// Shot performs a player's shot.
// @Description Произвести выстрел.
// @Tags battle
// @Accept  json
// @Produce  json
// @Param parameters body ShotReq true "JSON Body with parameters"
// @Success 200 {object} ShotResultResp
// @Failure 400 {string} string	"error"
// @Router /rest/v1/battle/shot [post]
func Shot(c *gin.Context) {
	var reqBody ShotReq

	if err := c.BindJSON(&reqBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	respBody, err := reqBody.Shot()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, respBody)
}

// ClearBattleField clears a sea battle game field.
// @Description Очистить поле для игры в морской бой.
// @Tags battle
// @Accept  json
// @Produce  json
// @Success 200 {string} string "message"
// @Router /rest/v1/battle/clear [post]
func ClearBattleField(c *gin.Context) {
	c.JSON(http.StatusOK, clearGameField())
}

// GetState returns game statistics.
// @Description Получить статистику по итогу игры.
// @Tags battle
// @Accept  json
// @Produce  json
// @Success 200 {object} GameState
// @Router /rest/v1/battle/state [get]
func GetState(c *gin.Context) {
	c.JSON(http.StatusOK, getGameState())
}
