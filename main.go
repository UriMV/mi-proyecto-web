package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Función personalizada para restar valores en el template
func sub(a, b int) int {
	return a - b
}

func main() {
	r := gin.Default()

	// Registrar funciones personalizadas en los templates
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"sub": sub,
	}).ParseGlob("templates/*"))

	r.SetHTMLTemplate(tmpl)

	// Cargar archivos estáticos
	r.Static("/static", "./static")

	// Rutas principales con breadcrumbs
	r.GET("/", func(c *gin.Context) {
		fmt.Println("Cargando página de inicio...")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":      "Inicio",
			"breadcrumb": []string{"Inicio"},
		})
	})

	r.GET("/acerca", func(c *gin.Context) {
		fmt.Println("Cargando página Acerca de...")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":      "Acerca de",
			"breadcrumb": []string{"Inicio", "Acerca de"},
		})
	})

	r.GET("/contacto", func(c *gin.Context) {
		fmt.Println("Cargando página de Contacto...")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":      "Contacto",
			"breadcrumb": []string{"Inicio", "Contacto"},
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"status":     404,
			"message":    "Página no encontrada",
			"breadcrumb": []string{"Inicio", "Error 404"},
		})
	})

	// Ruta para simular un error interno
	r.GET("/error500", func(c *gin.Context) {
		panic("Error interno simulado") // Forzar un error 500
	})
	

	// Iniciar servidor en el puerto 8080
	r.Run(":8080")
}
