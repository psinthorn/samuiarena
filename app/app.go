package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/psinthorn/samuiarena/configs"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	port := configs.Server.PortRunning("8091")
	// box := packr.NewBox("./views")
	router.LoadHTMLGlob("views/*/*.html")
	router.Static("/assets", "./assets")
	// router.Static("/views", "./views")
	urlsMapping()
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("ListenAndServe:"+port, err)
	}
}
