package main

import (
	"exporter-demo/config"
	"exporter-demo/handler"
	"exporter-demo/metric"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}

func main() {
	cfg, err := config.New()
	if err != nil {
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "default"
		zap.S().Warnf("unknow hostname, use default replace")
	}
	metric.Root = metric.NewCounter("myapp_processed_ops_total", "The total number of processed events", map[string]string{
		"app":      "demo",
		"hostname": hostname,
	})

	r := gin.Default()
	r.GET("/", handler.Root)
	r.GET("/metrics", handler.Metrics)

	if err := r.Run(cfg.Addr); err != nil {
		zap.S().Errorf("listen %q failed, err: %v", cfg.Addr, err)
		return
	}
}
