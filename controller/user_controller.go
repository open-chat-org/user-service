package controller

import (
	"strconv"
	"user-service/dto"
	"user-service/service"

	"github.com/gin-gonic/gin"
)

func NewController() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/api/v1.0/user/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": dto.FAILURE,
				"data":    nil,
			})
		}

		user, err := service.GetUser(id)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": dto.FAILURE,
				"data":    nil,
			})
		}

		c.JSON(200, gin.H{
			"code":    200,
			"message": dto.SUCCESSFUL,
			"data":    user,
		})
	})

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
