package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Title": "Homepage",
	})
}

func Person(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"First name": "",
		"Last name":  "",
		"Age":        "",
		"Weight": "",

	})
}
