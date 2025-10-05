package routes

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupReoutes(r *gin.Engine) {

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/:page", func(ctx *gin.Context) {
		page := ctx.Param("page")
		if !strings.HasSuffix(page, ".html") {
			page += ".html"
		}

		if _, err := os.Stat("templates" + page); err == nil {
			ctx.HTML(http.StatusOK, page, nil)
		} else {
			ctx.HTML(http.StatusNotFound, "404.html", nil)
		}
	})

}
