package main

import (
	"Gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupReoutes(r)

	r.Run(":4040")
}
