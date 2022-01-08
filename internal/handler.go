package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInterval() *gin.Engine {
	router := gin.Default()

	router.GET("/:from/:to", func(c *gin.Context) {
		from := c.Param("from")
		to := c.Param("to")

		c.String(http.StatusOK, from, to)
	})
	return router
}
