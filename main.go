package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func setupRouter(productionMode, jsonLoggingEnabled bool) *gin.Engine {
	var router *gin.Engine
	if productionMode {
		gin.SetMode(gin.ReleaseMode)
		if jsonLoggingEnabled {
			router = gin.New()
			router.Use(gin.Recovery())
			router.Use(gin.LoggerWithFormatter(JSONLogger))
		} else {
			router = gin.Default()
		}
	} else {
		router = gin.Default()
	}

	router.GET("/health", healthHandler)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/", mainHandler)

	return router
}

func main() {

	productionMode := strings.ToLower(os.Getenv("BLACKBIRD_ENV")) == "production"

	jsonLogging := os.Getenv("BLACKBIRD_JSON_LOGGING")
	jsonLoggingEnabled := strings.ToLower(jsonLogging) == "true"

	router := setupRouter(productionMode, jsonLoggingEnabled)

	http.ListenAndServe(":8080", router)
}
