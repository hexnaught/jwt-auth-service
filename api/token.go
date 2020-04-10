package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func validateTokenHandler(c *gin.Context) {

	authHeaderContent := c.Request.Header.Get("Authorization")
	authHeaderSplitContent := strings.Split(authHeaderContent, "Bearer")

	if len(authHeaderSplitContent) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if validateToken(strings.TrimSpace(authHeaderSplitContent[1])) != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Status(http.StatusOK)
}
