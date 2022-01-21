package internal

import (
	"FiboApp/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInterval() *gin.Engine {
	router := gin.Default()

	router.GET("/:from/:to", func(c *gin.Context) {

		from, to, err := pkg.Validate(c.Param("from"), c.Param("to"))
		if err == "" {
			response := pkg.JSON(mc.CheckMemcached(from, to), from)
			c.JSON(http.StatusOK, response)
		} else {
			c.String(http.StatusBadRequest, err)
		}
	})
	return router
}
