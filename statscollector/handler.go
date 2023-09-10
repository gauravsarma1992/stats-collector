package statscollector

import (
	"fmt"
	"net/http"

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
		metrics  []*Metric
		markdown string
		err      error
	)
	if metrics, err = server.metricStore.List(); err != nil {
		gorestapi.RequestBodyClientErrorHandler(c, err)
		return
	}
	markdown = "<html><head></head><body>default_total 10<br>"
	for _, metric := range metrics {
		markdown += fmt.Sprintf("%s_total %f<br>", metric.Name, metric.Value)
	}
	markdown += "<br></body></html>"
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(markdown))
	return
}
