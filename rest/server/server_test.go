package server

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestBuildServer(t *testing.T) {
	server := BuildServer()
	AddRoute(server, "GET", "hello", func(c *gin.Context) {
		c.String(200, "OK")
	})
	RunOnPort(server, 8081)
}
