package internal

import (
	"FiboApp/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetInterval() *gin.Engine {
	router := gin.Default()

	router.GET("/:from/:to", func(c *gin.Context) {
		from, err := strconv.Atoi(c.Param("from"))
		if err != nil {
			c.String(http.StatusBadRequest, "Use positive number for FROM")
		}
		to, err := strconv.Atoi(c.Param("to"))
		if err != nil {
			c.String(http.StatusBadRequest, "Use positive number for TO")
		}
		er := pkg.Validate(from, to)
		if er == "" {
			c.String(http.StatusOK, GenerateNum(from, to))
		} else {
			c.String(http.StatusBadRequest, er)
		}
	})
	return router
}
