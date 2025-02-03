package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Credenciales en memoria
var validUser = "admin"
var validPass = "1234"

func main() {
	r := gin.Default()

	// Cargar templates
	tmpl := template.Must(template.New("").ParseGlob("templates/*"))
	r.SetHTMLTemplate(tmpl)

	// Archivos estáticos
	r.Static("/static", "./static")

	// Página de login
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// Manejo del login
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == validUser && password == validPass {
			// Redirigir a la página de inicio tras el login exitoso
			c.Redirect(http.StatusSeeOther, "/")
		} else {
			// Mostrar mensaje de error
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"error": "Usuario o contraseña incorrectos",
			})
		}
	})

	// Ruta principal
	r.GET("/", func(c *gin.Context) {
		fmt.Println("Cargando página de inicio...")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":      "Inicio",
			"breadcrumb": []string{"Inicio"},
		})
	})

	// Iniciar servidor
	r.Run(":8080")
}
