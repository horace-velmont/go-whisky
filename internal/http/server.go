package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	engine *gin.Engine
	server *http.Server
}

type ServerConfig struct {
	Port string
}

func NewServer(
	config ServerConfig,
	routes []Routes,
) (*Server, error) {
	engine := gin.Default()
	server := &Server{
		engine: engine,
		server: &http.Server{
			Addr:    ":" + config.Port,
			Handler: engine,
		},
	}

	for _, route := range routes {
		route.Register(server.engine.Group(route.Path()))
	}

	return server, nil
}

func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

type Router interface {
	gin.IRouter
}

type Routes interface {
	Path() string
	Register(router Router)
}

type Middleware interface {
	Handler() gin.HandlerFunc
}
