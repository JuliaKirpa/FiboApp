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
			c.String(http.StatusBadRequest, "Use positive integer value for FROM\n")
		}
		to, err := strconv.Atoi(c.Param("to"))
		if err != nil {
			c.String(http.StatusBadRequest, "Use positive integer value for TO\n")
		}
		er := pkg.Validate(from, to)
		if er == "" {
			response := pkg.JSON(mc.CheckMemcached(from, to), from)
			c.JSON(http.StatusOK, response)
		} else {
			c.String(http.StatusBadRequest, er)
		}
	})
	return router
}
