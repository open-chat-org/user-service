package controller

import (
	"fmt"
	"strconv"
	"user-service/dto"
	"user-service/repository"
	"user-service/service"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func newController(driver neo4j.Driver) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	userService := service.UserNeo4jService{
		UserRepository: repository.UserNeo4jRepository{
			Drive: driver,
		},
	}

	router.GET("/api/v1.0/user/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Print(err.Error())
			c.JSON(400, dto.FailureResponse())
			return
		}

		user, err := userService.GetUser(id)
		if err != nil {
			fmt.Print(err.Error())
			c.JSON(400, dto.FailureResponse())
			return
		}

		c.JSON(200, dto.SuccessResponse(user))
	})

	router.PUT("/api/v0.1/user", func(c *gin.Context) {})

	router.DELETE("/api/v1.0/user/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Print(err.Error())
			c.JSON(400, dto.FailureResponse())
			return
		}

		err = userService.DeleteUser(id)
		if err != nil {
			fmt.Print(err.Error())
			c.JSON(400, dto.FailureResponse())
			return
		}

		c.JSON(200, dto.SuccessResponse(nil))
	})

	router.GET("/api/v0.1/user/list-friends", func(c *gin.Context) {})

	router.POST("/api/v0.1/user/invite-friend", func(c *gin.Context) {})

	router.PUT("/api/v0.1/user/accept-friend", func(c *gin.Context) {})

	return router
}

func InitGin(driver neo4j.Driver) {
	router := newController(driver)

	router.Run(":8080")
}
