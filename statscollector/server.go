package statscollector

import (
	"log"
	"os"

	gorestapi "github.com/gauravsarma1992/go-rest-api/gorestapi"
	"github.com/gin-gonic/gin"
)

type (
	Server struct {
		apiEngine *gorestapi.Server
	}
)

func NewServer() (server *Server, err error) {
	server = &Server{}
	if server.apiEngine, err = gorestapi.New(nil); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	if err = server.Setup(); err != nil {
		return
	}
	return
}

func successHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}

func (server *Server) Setup() (err error) {
	server.apiEngine.AddRoute(gorestapi.Route{"/api/success", "GET", successHandler, false})
	return
}

func (server *Server) Run() (err error) {
	if err = server.apiEngine.Run(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	return
}
