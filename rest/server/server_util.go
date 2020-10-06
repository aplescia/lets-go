package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//BuildServer returns a default gin Server Engine.
func BuildServer() *gin.Engine {
	return gin.Default()
}

//RunOnPort runs the specified gin Engine on the specified port.
func RunOnPort(engine *gin.Engine, port int){
	err := engine.Run(fmt.Sprintf(":%d",port))
	if err != nil {
		panic(err)
	}
}

//Add route will add a handler func under the specified path responding to the specified http method.
func AddRoute(server *gin.Engine, method string, path string, handler gin.HandlerFunc){
	switch method {
	case "GET":
		server.GET(path, handler)
	case "PUT":
		server.PUT(path, handler)
	case "POST":
		server.POST(path,handler)
	case "DELETE":
		server.DELETE(path,handler)
	case "PATCH":
		server.PATCH(path, handler)
	case "OPTIONS":
		server.OPTIONS(path,handler)
	case "HEAD":
		server.HEAD(path,handler)
	default:
		server.GET(path,handler)
	}
}
