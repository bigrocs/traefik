package router

import (
	"github.com/gin-gonic/gin"
	h "github.com/bigrocs/traefik/user-api/hanlder"
)

// Handler
func Handler(r *gin.Engine) {
	health := &h.Health{}
	r.GET("/health", func(c *gin.Context) {
		health.Health(c)
	})
}
