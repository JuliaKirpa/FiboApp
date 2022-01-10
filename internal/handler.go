package internal

import (
	"FiboApp/pkg"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetInterval() *gin.Engine {
	router := gin.Default()
	mc := memcache.New("memcached:11211")

	router.GET("/:from/:to", func(c *gin.Context) {
		from, err := strconv.Atoi(c.Param("from"))
		if err != nil {
			c.String(http.StatusBadRequest, "Use positive integer value for FROM")
		}
		to, err := strconv.Atoi(c.Param("to"))
		if err != nil {
			c.String(http.StatusBadRequest, "Use positive integer value for TO")
		}
		er := pkg.Validate(from, to)
		if er == "" {
			var response = pkg.Prepare(GenerateNum(from, to, mc), from)
			c.String(http.StatusOK, response)
		} else {
			c.String(http.StatusBadRequest, er)
		}
	})
	return router
}
