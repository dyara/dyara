package middlewares

import (
	"github.com/dyara/dyara/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

const urlPrefix = "/"

var staticFileServer = static.Setup()
var setupFileServer = http.StripPrefix(urlPrefix, http.FileServer(staticFileServer))

func SetupCheck(c *gin.Context) {
	setupRequired := c.GetBool("setup-required")

	if !setupRequired {
		// TODO load required config values
		// TODO validate dyara has required config
		// TODO if validate fail -> load setup SPA
		if staticFileServer.Exists(urlPrefix, c.Request.URL.Path) {
			setupFileServer.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		} else {
			c.Request.URL.Path = "/"
			setupFileServer.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	} else {
		// TODO else -> proceed
		c.Next()
	}
}
