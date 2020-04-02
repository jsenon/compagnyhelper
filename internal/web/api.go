package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jsenon/compagnyhelper/internal/link"
	"github.com/opentracing/opentracing-go"
)

// In cli compagnyhelper get link -n dev, output will be
// Grafana \n Kibana \n Prometheus \n
//
// In cli compagnyhelper get link --all, output will be
// Grafana dev \n Kibana dev \n Grafana prod \n
func retrieveLinks(c *gin.Context) {
	span, ctxChild := opentracing.StartSpanFromContext(c.Request.Context(), "(*compagnyhelper).web.retrieveLinks")
	defer span.Finish()

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	links, err := link.RetrieveAll(ctxChild, request.Env)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, links)
}

// In cli compagnyhelper get link grafana -n dev, output will be
// Grafana |  http://grafana.com \n
//
// In cli compagnyhelper describe link grafana -n dev
// Grafana | Your dashboard for your metrics | development | http://grafana.com \n
func retrievedescribeLink(c *gin.Context) {
	span, ctxChild := opentracing.StartSpanFromContext(c.Request.Context(), "(*compagnyhelper).web.retrievedescribeLink")
	defer span.Finish()

	name := c.Param("name")

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link, err := link.Retrieve(ctxChild, request.Env, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, link)
}

// In cli compagnyhelper open link grafana -n dev, output will open a browser with url http://grafana.com
func openLink(c *gin.Context) {
	span, ctxChild := opentracing.StartSpanFromContext(c.Request.Context(), "(*compagnyhelper).web.openLink")
	defer span.Finish()

	name := c.Param("name")

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link, err := link.Retrieve(ctxChild, request.Env, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, link.Desc.Link)
}
