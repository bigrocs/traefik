package router

import (
	h "github.com/bigrocs/traefik/user-api/handler"
	"github.com/gin-gonic/gin"
)

// Handler
func Handler(r *gin.Engine) {
	health := &h.Health{}
	r.GET("/health", func(c *gin.Context) {
		health.Health(c)
	})
}
