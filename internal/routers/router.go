package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", Pong) // v1/ping
	}
	v2 := r.Group("v2")
	{
		v2.GET("/ping/:name", Pong)
	}

	return r
}

func Pong(c *gin.Context) {
	// name := c.Param("name")
	name := c.DefaultQuery("name", "Backend-Go-Core")
	// c.ShouldBindJSON()
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + name,
		"uid":     uid,
		"users":   []string{"you", "me"},
	})
}
