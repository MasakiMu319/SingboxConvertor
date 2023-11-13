package api

import (
	"github.com/gin-gonic/gin"
)

// GetFrontend show index page.
func GetFrontend(c *gin.Context) {
	//c.File("./web/frontend.html")
	c.File("/server/web/frontend.html")
}
