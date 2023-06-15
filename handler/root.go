package handler

import (
	"exporter-demo/metric"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
	metric.Root.Inc()
	// zap.S().Infof("root metric desc: %s", metric.Root.Desc().String())
	c.JSON(http.StatusOK, gin.H{
		"result":  http.StatusOK,
		"message": "request success",
	})
}
