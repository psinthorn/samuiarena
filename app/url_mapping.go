package app

import "github.com/psinthorn/samuiarena/controllers"

func urlsMapping() {
	router.GET("/", controllers.GetIndex)
}
