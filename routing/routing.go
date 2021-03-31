package routing

import (
	cfg "../config"
	log "../logger"

	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func Setup() {
	log.InfoLogger.Println("Initialize gin engine...")
	router = gin.Default()

	log.InfoLogger.Println("Setting GIN to %v mode", cfg.ServerSetting.RunMode)
	if cfg.ServerSetting.RunMode == "debug"{
		gin.SetMode(gin.DebugMode)
	}else if cfg.ServerSetting.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}else{
		log.FatalLogger.Println("GIN mode should be release or debug!")
		panic("")
	}

	log.InfoLogger.Println("Loading HTML templates...")
	router.LoadHTMLGlob("templates/*")

	initRoutes()
}

func initRoutes() {
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			})
	})
}

func Run() {
	router.Run(cfg.ServerSetting.HttpAddress + ":" + cfg.ServerSetting.HttpPort)
}
