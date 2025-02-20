package main

import (
	"github.com/HaikalRFadhilahh/api-mrt-go-bwa/modules/station"
	"github.com/gin-gonic/gin"
)

func main() {
	InitiateRouter()
}

func InitiateRouter() {
	// Init Router Gin
	router := gin.Default()

	// Group API Endpoint
	api := router.Group("/v1/api")

	// Parent Routing
	station.Initiate(api)

	// Run Gin Server
	router.Run(":8080")
}
