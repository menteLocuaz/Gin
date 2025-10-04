package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupReoutes(r *gin.Engine) {

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Title":   "Mi aplicacion",
			"Heading": "Hola mundo",
			"Message": "Bienvenid a mi aplicacion web con gin y platilla HTML",
		})
	})

	r.Static("/static", "./static")

}
