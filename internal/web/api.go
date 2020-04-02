package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func retrieveLinks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "restrie!"})
}

func describeLink(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "describe!"})
}

func openLink(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "open!"})
}
