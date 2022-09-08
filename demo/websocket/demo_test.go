package websocket

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestDemo(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", ping)
	r.Run(":12345")
}
