package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuario []Usuario

func SetupReoutes(r *gin.Engine) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hola mundo")
	})

	r.GET("/saludo/:nombre", func(ctx *gin.Context) {
		nombre := ctx.Param("nombre")
		ctx.String(http.StatusOK, "Hola , - %s ", nombre)
	})

	r.POST("/usuarios", func(ctx *gin.Context) {
		var nuevoUsuario Usuario
		if err := ctx.BindJSON(&nuevoUsuario); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Error al decodificar el JSON",
			})
			return
		}

		if nuevoUsuario.Nombre == "" || nuevoUsuario.Email == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "El email y nombre son datos requeridos",
			})
			return
		}

		usuario = append(usuario, nuevoUsuario)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Usuario registrado",
			"datos":   usuario,
		})
	})

}
