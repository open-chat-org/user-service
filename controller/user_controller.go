package controller

import "github.com/gin-gonic/gin"

func NewController() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/api/v0.1/user", func(c *gin.Context) {})

	router.PUT("/api/v0.1/user", func(c *gin.Context) {})

	router.DELETE("/api/v0.1/user", func(c *gin.Context) {})

	router.GET("/api/v0.1/user/list-friends", func(c *gin.Context) {})

	router.POST("/api/v0.1/user/invite-friend", func(c *gin.Context) {})

	router.PUT("/api/v0.1/user/accept-friend", func(c *gin.Context) {})

	router.GET("/internal/api/v0.1/user", func(c *gin.Context) {})

	router.PUT("/internal/api/v0.1/user", func(c *gin.Context) {})

	return router
}

func Init() {
	router := NewController()

	router.Run(":8080")
}
