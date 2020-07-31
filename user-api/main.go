package main

import (
	// 公共引入

	"github.com/gin-gonic/gin"

	"github.com/bigrocs/traefik/user-api/router"
)

func main() {
	r := gin.Default()
	router.Handler(r) // 路由实现
	r.Run(":1080") // listen and serve on 0.0.0.0:8080
}
