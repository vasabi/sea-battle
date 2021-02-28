package main

import (
	"sea-battle/docs"
	"sea-battle/internal/app"
	"sea-battle/internal/pkg/utils"
)

// @title Sea battle rest API
// @version 1.0
// @description Морской бой

// @contact.name SeaBattle
// @contact.url https://github.com/vasabi/sea-battle
// @contact.email vasabijaj@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath
// @schemes http https

func main() {
	docs.SwaggerInfo.Host = utils.GetSwaggerInfoHost()
	app.Run()
}
