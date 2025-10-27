package routes

import "github.com/gin-gonic/gin"

func SetupRoute() *gin.Engine {
	r:=gin.Default()
	api := r.Group("/api")

	return r
}