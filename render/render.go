package render

import (
	"github.com/gin-gonic/gin"
	"lemonadee/internal"
	"net/http"
)

func BadRequest(c *gin.Context) {
	message := "Bad request"
	c.JSON(http.StatusBadRequest, gin.H{
		"error": message,
	})
}

func InternalServerError(c *gin.Context) {
	message := "InternalServerError"
	c.JSON(http.StatusBadRequest, gin.H{
		"error": message,
	})
}

func Error(c *gin.Context, err *internal.Error) {
	c.JSON(err.StatusCode, gin.H{
		"error": err,
	})
}
