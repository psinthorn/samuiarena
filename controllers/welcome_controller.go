package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", nil)
}
