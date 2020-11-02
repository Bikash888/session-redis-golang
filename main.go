package main

import (
	"session-redis/src/controller"
	"session-redis/src/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.POST("/login", controller.Login)
	r.GET("/logout", controller.Logout)

	auth := r.Group("/auth")
	auth.Use(middleware.Authentication())
	{
		auth.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Everything is ok",
			})
		})

	}
	r.Run(":8000")
}
