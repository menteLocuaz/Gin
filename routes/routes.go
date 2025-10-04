package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupReoutes(r *gin.Engine) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hola mundo")
	})

	r.GET("/saludo/:nombre", func(ctx *gin.Context) {
		nombre := ctx.Param("nombre")
		ctx.String(http.StatusOK, "Hola , - %s ", nombre)
	})

}
