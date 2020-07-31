package main

import (
	// 公共引入

	"github.com/gin-gonic/gin"

	h "github.com/bigrocs/traefik/user-api/handler"
)

func main() {
	r := gin.Default()
	health := &h.Health{}
	r.GET("/health", func(c *gin.Context) {
		health.Health(c)
	})
	r.Run(":1080") // listen and serve on 0.0.0.0:8080
}
