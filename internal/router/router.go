package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()

	r.Run(":8080")
}
