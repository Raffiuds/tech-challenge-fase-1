package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRouter retorna um novo router.
func NewRouter(router *gin.Engine) *gin.Engine {

	router.GET("/", index)

	return router
}

// Index rota raiz
func index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
