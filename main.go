package main

import (
	"fmt"
	"log"
	"net/http"
	"productos-crud/controllers"
	"productos-crud/database"
)

func main() {
	// Inicializar base de datos
	database.InitDB()
	defer database.CloseDB()
	
	// Crear directorio uploads
	database.CreateUploadsDir()

	// Servir archivos estáticos
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	// Rutas principales
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/productos", http.StatusSeeOther)
			return
		}
		http.NotFound(w, r)
	})

	// Instanciar controladores
	productoController := &controllers.ProductoController{}
	categoriaController := &controllers.CategoriaController{}

	// Rutas de productos
	http.HandleFunc("/productos", productoController.Index)
	http.HandleFunc("/producto/", productoController.HandleRoutes)

	// Rutas de categorías
	http.HandleFunc("/categorias", categoriaController.HandleRoutes)

	// API endpoints
	http.HandleFunc("/api/productos", productoController.APIIndex)
	http.HandleFunc("/api/categorias", categoriaController.APIIndex)

	fmt.Println("Servidor iniciado en http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
