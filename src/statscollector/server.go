package statscollector

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	Server struct {
		Config    *ServerConfig
		apiEngine *gin.Engine
	}

	ServerConfig struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
)

func NewServer() (server *Server, err error) {
	server = &Server{}
	if server.apiEngine, err = server.DefaultEngineConfig(); err != nil {
		return
	}
	if err = server.SetRoutes(); err != nil {
		return
	}
	return
}

func (server *Server) DefaultEngineConfig() (apiEngine *gin.Engine, err error) {
	apiEngine = gin.Default()
	server.Config = &ServerConfig{
		Host: "127.0.0.1",
		Port: "8090",
	}
	return
}

func (server *Server) SetRoutes() (err error) {
	server.apiEngine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return
}

func (server *Server) Run() (err error) {
	server.apiEngine.Run(fmt.Sprintf("%s:%s", server.Config.Host, server.Config.Port))
	return
}
