package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jsenon/compagnyhelper/internal/link"
)

func retrieveLinks(c *gin.Context) {
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	links, err := link.RetrieveAll(c.Request.Context(), request.Env)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, links)
}

func retrieveLink(c *gin.Context) {
	name := c.Param("name")

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link, err := link.Retrieve(c.Request.Context(), request.Env, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, link)
}

func describeLink(c *gin.Context) {
	name := c.Param("name")

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link, err := link.DescribeLink(c.Request.Context(), request.Env, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, link)
}

func openLink(c *gin.Context) {
	name := c.Param("name")

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link, err := link.DescribeLink(c.Request.Context(), request.Env, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, link.Desc.Link)
}
