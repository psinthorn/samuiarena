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
	router.LoadHTMLGlob("views/*/*.html")
	router.Static("/assets", "./assets")
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("ListenAndServe:"+port, err)
	}
}
