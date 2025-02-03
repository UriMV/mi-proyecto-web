package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controlador para la página de inicio
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":      "Inicio",
		"breadcrumb": []string{"Inicio"},
	})
}

// Controlador para la página "Acerca de"
func Acerca(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":      "Acerca de",
		"breadcrumb": []string{"Inicio", "Acerca de"},
	})
}

// Controlador para la página de contacto
func Contacto(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":      "Contacto",
		"breadcrumb": []string{"Inicio", "Contacto"},
	})
}

// Controlador para la página de error 500
func Error500(c *gin.Context) {
	c.HTML(http.StatusInternalServerError, "error.html", gin.H{
		"status":     500,
		"message":    "Error interno del servidor",
		"breadcrumb": []string{"Inicio", "Error 500"},
	})
}

// Controlador para las páginas no encontradas
func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error.html", gin.H{
		"status":     404,
		"message":    "Página no encontrada",
		"breadcrumb": []string{"Inicio", "Error 404"},
	})
}
