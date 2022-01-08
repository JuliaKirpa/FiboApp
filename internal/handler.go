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
			c.String(http.StatusBadRequest, "Use uint value for FROM")
		}
		to, err := strconv.Atoi(c.Param("to"))
		if err != nil {
			c.String(http.StatusBadRequest, "Use uint value for TO")
		}

		er := pkg.Validate(from, to)
		var response = pkg.Prepare(GenerateNum(from, to), from)

		if er == "" {
			c.String(http.StatusOK, response)
		} else {
			c.String(http.StatusBadRequest, er)
		}
	})
	return router
}
