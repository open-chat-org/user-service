package dto

import "github.com/gin-gonic/gin"

const (
	FAILURE    = "FAILURE"
	SUCCESSFUL = "SUCCESSFUL"
)

func SuccessResponse(data interface{}) gin.H {
	return gin.H{
		"code":    200,
		"message": SUCCESSFUL,
		"data":    data,
	}
}

func FailureResponse() gin.H {
	return gin.H{
		"code":    400,
		"message": FAILURE,
		"data":    nil,
	}
}
