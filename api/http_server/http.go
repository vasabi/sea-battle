package http_server

import (
	"net/http"
	"sea-battle/internal/pkg/utils"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

// RunServer starts the http server.
func RunServer() {
	if utils.GetDebugModeON() {
		gin.SetMode(gin.DebugMode)
		log.SetLevel(log.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
		log.SetLevel(log.InfoLevel)
	}

	router := gin.New()
	router.RedirectTrailingSlash = true

	router.Use(
		ginlogrus.Logger(log.StandardLogger()),
		gin.Recovery(),
	)

	// Default response for not found pages.
	router.NoRoute(func(c *gin.Context) {
		var allowableRoute = make(map[string][]string)

		log.Errorf("page not found: %v", c.Request.URL)

		for _, route := range router.Routes() {
			allowableRoute[route.Method] = append(allowableRoute[route.Method], route.Path)
		}

		c.AbortWithStatusJSON(http.StatusNotFound, struct {
			Routes map[string][]string `json:"allowable_routes"`
		}{
			allowableRoute,
		})
	})

	log.Info("init HTTP server at :" + utils.GetHTTPPort())
	initRouter(router)

	err := router.Run(":" + utils.GetHTTPPort())
	if err != nil {
		log.Fatal("cannot start server")
	}
}
