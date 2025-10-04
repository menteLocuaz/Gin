package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		// ctx.String(200, "Hello, world")
		ctx.JSON(200, gin.H{
			"message": "Hola mundo",
		})
	})

	r.Run(":4040")
}
