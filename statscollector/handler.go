package statscollector

import (
	"github.com/gauravsarma1992/go-rest-api/gorestapi"
	"github.com/gin-gonic/gin"
)

func (server *Server) MetricsAddHandler(c *gin.Context) {
	var (
		metric Metric
		err    error
	)
	if err = c.ShouldBindJSON(&metric); err != nil {
		gorestapi.RequestBodyClientErrorHandler(c, err)
		return
	}
	if err = server.metricStore.Add(&metric); err != nil {
		gorestapi.ResourceNotFoundHandler(c, metric.Name)
		return
	}
	return
}

func (server *Server) MetricsListHandler(c *gin.Context) {
	var (
		metrics []*Metric
		err     error
	)
	if metrics, err = server.metricStore.List(); err != nil {
		gorestapi.RequestBodyClientErrorHandler(c, err)
		return
	}
	c.JSON(200, gin.H{
		"metrics": metrics,
	})
	return
}
